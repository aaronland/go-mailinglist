package http

// CSRF crumbs are handled by go-http-crumb middleware
// Bootstrap stuff is handled by go-http-bootstrap middleware
// see cmd/subscriptiond/main.go for details

import (
	"fmt"
	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/eventlog"
	"github.com/aaronland/go-mailinglist/invitation"
	"github.com/aaronland/go-mailinglist/message"
	"github.com/aaronland/gomail"
	"html/template"
	"log"
	gohttp "net/http"
	"net/mail"
	"net/url"
	"time"
)

type InviteRequestTemplateVars struct {
	SiteName string
	Paths    *mailinglist.PathConfig
	Error    error
}

type InviteRequestHandlerOptions struct {
	Config        *mailinglist.MailingListConfig
	Templates     *template.Template
	Subscriptions database.SubscriptionsDatabase
	Invitations   database.InvitationsDatabase
	EventLogs     database.EventLogsDatabase
	Sender        gomail.Sender
}

type InviteRequestEmailTemplateVars struct {
	Invites  []*invitation.Invitation
	SiteName string
	Paths    *mailinglist.PathConfig
}

func InviteRequestHandler(opts *InviteRequestHandlerOptions) (gohttp.Handler, error) {

	invite_t, err := LoadTemplate(opts.Templates, "invite_request")

	if err != nil {
		return nil, err
	}

	success_t, err := LoadTemplate(opts.Templates, "invite_request_success")

	if err != nil {
		return nil, err
	}

	email_t, err := LoadTemplate(opts.Templates, "invite_request_email")

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		vars := InviteRequestTemplateVars{
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
		}

		if !opts.Config.FeatureFlags.InviteRequest {

			app_err := NewApplicationError(nil, E_DISABLED_INVITE)
			vars.Error = app_err

			RenderTemplate(rsp, invite_t, vars)
			return
		}

		switch req.Method {

		case "GET":
			RenderTemplate(rsp, invite_t, vars)
			return

		case "POST":

			subs_db := opts.Subscriptions
			invites_db := opts.Invitations

			str_addr, err := sanitize.PostString(req, "address")

			if err != nil {

				app_err := NewApplicationError(err, E_INPUT_PARSE, "address")
				vars.Error = app_err

				RenderTemplate(rsp, invite_t, vars)
				return
			}

			if str_addr == "" {

				app_err := NewApplicationError(err, E_INPUT_MISSING, "address")
				vars.Error = app_err

				RenderTemplate(rsp, invite_t, vars)
				return
			}

			addr, err := mail.ParseAddress(str_addr)

			if err != nil {

				app_err := NewApplicationError(err, E_EMAIL_PARSE)
				vars.Error = app_err
				RenderTemplate(rsp, invite_t, vars)
				return
			}

			sub, err := subs_db.GetSubscriptionWithAddress(addr.Address)

			if err != nil {

				app_err := NewApplicationError(err, E_SUBSCRIPTION_RETRIEVE)
				vars.Error = app_err
				RenderTemplate(rsp, invite_t, vars)
				return
			}

			if !sub.IsEnabled() {
				app_err := NewApplicationError(err, E_SUBSCRIPTION_DISABLED)
				vars.Error = app_err
				RenderTemplate(rsp, invite_t, vars)
				return
			}

			// START all of this should go in a function

			invites := make([]*invitation.Invitation, 0)
			counts := make(map[string]int)

			invites_cb := func(invite *invitation.Invitation) error {

				t := time.Unix(invite.Created, 0)
				yyyymm := t.Format("200601")

				count, ok := counts[yyyymm]

				if !ok {
					count = 0
				}

				counts[yyyymm] = count + 1

				if invite.IsAvailable() {
					invites = append(invites, invite)
				}

				return nil
			}

			err = invites_db.ListInvitationsWithInviter(req.Context(), invites_cb, sub)

			if err != nil {
				app_err := NewApplicationError(err, E_INVITATION_LIST)
				vars.Error = app_err
				RenderTemplate(rsp, invite_t, vars)
				return
			}

			now := time.Now()
			yyyymm := now.Format("200601")

			count, ok := counts[yyyymm]

			if !ok {
				count = 0
			}

			max_codes := 2 - count

			if max_codes <= 0 {
				app_err := NewApplicationError(nil, E_INVITATION_MAX)
				vars.Error = app_err
				RenderTemplate(rsp, invite_t, vars)
				return
			}

			for len(invites) < max_codes {

				invite, err := invitation.NewInvitation(sub)

				if err != nil {
					app_err := NewApplicationError(err, E_INVITATION_NEW)
					vars.Error = app_err
					RenderTemplate(rsp, invite_t, vars)
					return
				}

				err = invites_db.AddInvitation(invite)

				if err != nil {
					app_err := NewApplicationError(err, E_INVITATION_ADD)
					vars.Error = app_err
					RenderTemplate(rsp, invite_t, vars)
					return
				}

				invites = append(invites, invite)
			}

			// END all of this should go in a function

			for _, i := range invites {
				log.Println("CODE ", i.Code)
			}

			invite_event_params := url.Values{}
			invite_event_params.Set("remote_addr", req.RemoteAddr)

			for _, i := range invites {
				invite_event_params.Set("invite_code", i.Code)
			}

			invite_event_message := invite_event_params.Encode()

			invite_event := &eventlog.EventLog{
				Address: addr.Address,
				Created: time.Now().Unix(),
				Event:   eventlog.EVENTLOG_INVITE_REQUEST_EVENT,
				Message: invite_event_message,
			}

			invite_event_err := opts.EventLogs.AddEventLog(invite_event)

			if invite_event_err != nil {
				log.Println(invite_event_err)
			}

			email_vars := InviteRequestEmailTemplateVars{
				Invites:  invites,
				SiteName: opts.Config.Name,
				Paths:    opts.Config.Paths,
			}

			msg, err := message.NewMessageFromHTMLTemplate(email_t, email_vars)

			if err != nil {
				app_err := NewApplicationError(err, E_EMAIL_CREATE)
				vars.Error = app_err
				RenderTemplate(rsp, invite_t, vars)
				return
			}

			from_addr, _ := mail.ParseAddress(opts.Config.Sender)
			to_addr, _ := mail.ParseAddress(sub.Address)

			subject := fmt.Sprintf("Your subscription to the %s mailing list", opts.Config.Name)

			msg_opts := &message.SendMessageOptions{
				Sender:  opts.Sender,
				Subject: subject,
				From:    from_addr,
				To:      to_addr,
			}

			send_err := message.SendMessage(msg, msg_opts)

			send_event_params := url.Values{}
			send_event_params.Set("remote_addr", req.RemoteAddr)

			for _, i := range invites {
				send_event_params.Set("invite_code", i.Code)
			}

			send_event_params.Set("action", "invite_request")

			send_event_id := eventlog.EVENTLOG_SEND_OK_EVENT

			if send_err != nil {
				send_event_id = eventlog.EVENTLOG_SEND_FAIL_EVENT
				send_event_params.Set("error", send_err.Error())
			}

			send_event_message := send_event_params.Encode()

			send_event := &eventlog.EventLog{
				Address: addr.Address,
				Created: time.Now().Unix(),
				Event:   send_event_id,
				Message: send_event_message,
			}

			send_event_err := opts.EventLogs.AddEventLog(send_event)

			if send_event_err != nil {
				log.Println(send_event_err)
			}

			if send_err != nil {
				app_err := NewApplicationError(send_err, E_EMAIL_SEND)
				vars.Error = app_err
				RenderTemplate(rsp, invite_t, vars)
				return
			}

			RenderTemplate(rsp, success_t, vars)
			return

		default:
			gohttp.Error(rsp, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
