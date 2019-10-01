package http

// CSRF crumbs are handled by go-http-crumb middleware
// Bootstrap stuff is handled by go-http-bootstrap middleware
// see cmd/subscriptiond/main.go for details

import (
	"errors"
	"fmt"
	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/confirmation"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/eventlog"
	"github.com/aaronland/go-mailinglist/message"
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/aaronland/gomail"
	"html/template"
	"log"
	gohttp "net/http"
	"net/mail"
	"net/url"
	"time"
)

type InviteCodeTemplateVars struct {
	SiteName string
	Paths    *mailinglist.PathConfig
	Error    error
}

type InviteAcceptTemplateVars struct {
	SiteName string
	Paths    *mailinglist.PathConfig
	Error    error
	Code     string
}

type InviteSubscribeTemplateVars struct {
	SiteName string
	Paths    *mailinglist.PathConfig
	Error    error
}

type InviteAcceptHandlerOptions struct {
	Config        *mailinglist.MailingListConfig
	Templates     *template.Template
	Subscriptions database.SubscriptionsDatabase
	Confirmations database.ConfirmationsDatabase
	Invitations   database.InvitationsDatabase
	EventLogs     database.EventLogsDatabase
	Sender        gomail.Sender
}

func InviteAcceptHandler(opts *InviteAcceptHandlerOptions) (gohttp.Handler, error) {

	code_t, err := LoadTemplate(opts.Templates, "invite_code")

	if err != nil {
		return nil, err
	}

	accept_t, err := LoadTemplate(opts.Templates, "invite_accept")

	if err != nil {
		return nil, err
	}

	success_t, err := LoadTemplate(opts.Templates, "subscribe_success")

	if err != nil {
		return nil, err
	}

	email_t, err := LoadTemplate(opts.Templates, "confirm_email")

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		code_vars := InviteCodeTemplateVars{
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
		}

		accept_vars := InviteAcceptTemplateVars{
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
		}

		if !opts.Config.FeatureFlags.InviteAccept {
			code_vars.Error = errors.New("Disabled")
			RenderTemplate(rsp, code_t, code_vars)
			return
		}

		invites_db := opts.Invitations

		switch req.Method {

		case "GET":

			code, err := sanitize.GetString(req, "code")

			if err != nil {
				code_vars.Error = err
				RenderTemplate(rsp, code_t, code_vars)
				return
			}

			if code == "" {
				RenderTemplate(rsp, code_t, code_vars)
				return
			}

			invite, err := invites_db.GetInvitationWithCode(code)

			if err != nil {
				code_vars.Error = err
				RenderTemplate(rsp, code_t, code_vars)
				return
			}

			if !invite.IsAvailable() {
				code_vars.Error = errors.New("Invitation unavailable")
				RenderTemplate(rsp, code_t, code_vars)
				return
			}

			// now we switch to accept_t

			accept_vars.Code = code

			RenderTemplate(rsp, accept_t, accept_vars)
			return

		case "POST":

			subs_db := opts.Subscriptions
			conf_db := opts.Confirmations

			code, err := sanitize.PostString(req, "code")

			if err != nil {
				code_vars.Error = err
				RenderTemplate(rsp, code_t, code_vars)
				return
			}

			if code == "" {
				code_vars.Error = errors.New("Missing code")
				RenderTemplate(rsp, code_t, code_vars)
				return
			}

			invite, err := invites_db.GetInvitationWithCode(code)

			if err != nil {
				code_vars.Error = err
				RenderTemplate(rsp, code_t, code_vars)
				return
			}

			if !invite.IsAvailable() {
				code_vars.Error = errors.New("Invitation unavailable")
				RenderTemplate(rsp, code_t, code_vars)
				return
			}

			// now we switch to accept_t

			accept_vars.Code = code

			str_addr, err := sanitize.PostString(req, "address")

			if err != nil {
				accept_vars.Error = err
				RenderTemplate(rsp, accept_t, accept_vars)
				return
			}

			confirmed, err := sanitize.PostString(req, "confirm")

			if err != nil {
				accept_vars.Error = err
				RenderTemplate(rsp, accept_t, accept_vars)
				return
			}

			if confirmed == "" {
				accept_vars.Error = errors.New("Unconfirmed")
				RenderTemplate(rsp, accept_t, accept_vars)
				return
			}

			if str_addr == "" {
				accept_vars.Error = errors.New("Empty address")
				RenderTemplate(rsp, accept_t, accept_vars)
				return
			}

			addr, err := mail.ParseAddress(str_addr)

			if err != nil {
				accept_vars.Error = err
				RenderTemplate(rsp, accept_t, accept_vars)
				return
			}

			sub, err := subs_db.GetSubscriptionWithAddress(addr.Address)

			// PLEASE MAKE ME A BETTER ERROR HANDLER...

			if sub != nil {
				accept_vars.Error = errors.New("Exists")
				RenderTemplate(rsp, accept_t, accept_vars)
				return
			}

			// START please reconcile me with http/subscribe.go

			sub, err = subscription.NewSubscription(addr.Address)

			if err != nil {
				accept_vars.Error = err
				RenderTemplate(rsp, accept_t, accept_vars)
				return
			}

			conf, err := confirmation.NewConfirmationForSubscription(sub, "subscribe")

			if err != nil {
				accept_vars.Error = err
				RenderTemplate(rsp, accept_t, accept_vars)
				return
			}

			err = subs_db.AddSubscription(sub)

			if err != nil {
				accept_vars.Error = err
				RenderTemplate(rsp, accept_t, accept_vars)
				return
			}

			err = conf_db.AddConfirmation(conf)

			if err != nil {

				go subs_db.RemoveSubscription(sub)

				accept_vars.Error = err
				RenderTemplate(rsp, accept_t, accept_vars)
				return
			}

			subscribe_event_params := url.Values{}
			subscribe_event_params.Set("remote_addr", req.RemoteAddr)
			subscribe_event_params.Set("confirmation_code", conf.Code)
			subscribe_event_params.Set("invitation_code", invite.Code)

			subscribe_event_message := subscribe_event_params.Encode()

			subscribe_event := &eventlog.EventLog{
				Address: addr.Address,
				Created: time.Now().Unix(),
				Event:   eventlog.EVENTLOG_SUBSCRIBE_EVENT,
				Message: subscribe_event_message,
			}

			subscribe_event_err := opts.EventLogs.AddEventLog(subscribe_event)

			if subscribe_event_err != nil {
				log.Println(subscribe_event_err)
			}

			email_vars := ConfirmationEmailTemplateVars{
				Code:     conf.Code,
				SiteName: opts.Config.Name,
				Paths:    opts.Config.Paths,
				Action:   "subscribe",
			}

			msg, err := message.NewMessageFromHTMLTemplate(email_t, email_vars)

			if err != nil {
				accept_vars.Error = err
				RenderTemplate(rsp, accept_t, accept_vars)
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
			send_event_params.Set("confirmation_code", conf.Code)
			send_event_params.Set("invitation_code", invite.Code)
			send_event_params.Set("action", conf.Action)

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
				accept_vars.Error = send_err
				RenderTemplate(rsp, accept_t, accept_vars)
				return
			}

			subscribe_vars := InviteSubscribeTemplateVars{
				SiteName: opts.Config.Name,
				Paths:    opts.Config.Paths,
			}

			RenderTemplate(rsp, success_t, subscribe_vars)
			return

			// END please reconcile me with http/subscribe.go

		default:
			gohttp.Error(rsp, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
