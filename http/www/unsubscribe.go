package www

// CSRF crumbs are handled by go-http-crumb middleware
// Bootstrap stuff is handled by go-http-bootstrap middleware
// see cmd/subscriptiond/main.go for details

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/mail"
	"net/url"
	"time"

	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist/v2"
	"github.com/aaronland/go-mailinglist/v2/confirmation"
	"github.com/aaronland/go-mailinglist/v2/database"
	"github.com/aaronland/go-mailinglist/v2/eventlog"
	"github.com/aaronland/go-mailinglist/v2/message"
	"github.com/aaronland/gomail/v2"
)

type UnsubscribeTemplateVars struct {
	SiteName string
	Paths    *mailinglist.PathConfig
	Error    error
}

type UnsubscribeHandlerOptions struct {
	Config        *mailinglist.MailingListConfig
	Templates     *template.Template
	Subscriptions database.SubscriptionsDatabase
	Confirmations database.ConfirmationsDatabase
	EventLogs     database.EventLogsDatabase
	Sender        gomail.Sender
}

func UnsubscribeHandler(opts *UnsubscribeHandlerOptions) (http.Handler, error) {

	unsubscribe_t, err := LoadTemplate(opts.Templates, "unsubscribe")

	if err != nil {
		return nil, err
	}

	success_t, err := LoadTemplate(opts.Templates, "unsubscribe_success")

	if err != nil {
		return nil, err
	}

	email_t, err := LoadTemplate(opts.Templates, "confirm_email")

	if err != nil {
		return nil, err
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		vars := UnsubscribeTemplateVars{
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
		}

		if !opts.Config.FeatureFlags.Unsubscribe {
			app_err := NewApplicationError(nil, E_DISABLED_UNSUBSCRIBE)
			vars.Error = app_err
			RenderTemplate(rsp, unsubscribe_t, vars)
			return
		}

		ctx := req.Context()

		switch req.Method {

		case "GET":

			RenderTemplate(rsp, unsubscribe_t, vars)
			return

		case "POST":

			subs_db := opts.Subscriptions
			conf_db := opts.Confirmations

			str_addr, err := sanitize.PostString(req, "address")

			if err != nil {
				app_err := NewApplicationError(err, E_INPUT_PARSE, "address")
				vars.Error = app_err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			addr, err := mail.ParseAddress(str_addr)

			if err != nil {
				app_err := NewApplicationError(err, E_EMAIL_PARSE)
				vars.Error = app_err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			sub, err := subs_db.GetSubscriptionWithAddress(ctx, addr.Address)

			if err != nil {

				if !database.IsNotExist(err) {
					app_err := NewApplicationError(err, E_SUBSCRIPTION_RETRIEVE)
					vars.Error = app_err
					RenderTemplate(rsp, unsubscribe_t, vars)
					return
				}

				app_err := NewApplicationError(err, E_SUBSCRIPTION_NOTFOUND)
				vars.Error = app_err

				vars.Error = app_err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			if !sub.IsEnabled() {
				app_err := NewApplicationError(err, E_SUBSCRIPTION_DISABLED)
				vars.Error = app_err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			conf, err := confirmation.NewConfirmationForSubscription(sub, "unsubscribe")

			if err != nil {
				app_err := NewApplicationError(err, E_CONFIRMATION_CREATE)
				vars.Error = app_err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			err = conf_db.AddConfirmation(ctx, conf)

			if err != nil {

				app_err := NewApplicationError(err, E_CONFIRMATION_ADD)
				vars.Error = app_err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			unsubscribe_event_params := url.Values{}
			unsubscribe_event_params.Set("remote_addr", req.RemoteAddr)
			unsubscribe_event_params.Set("confirmation_code", conf.Code)

			unsubscribe_event_message := unsubscribe_event_params.Encode()

			unsubscribe_event := &eventlog.EventLog{
				Address: addr.Address,
				Created: time.Now().Unix(),
				Event:   eventlog.EVENTLOG_UNSUBSCRIBE_EVENT,
				Message: unsubscribe_event_message,
			}

			unsubscribe_event_err := opts.EventLogs.AddEventLog(ctx, unsubscribe_event)

			if unsubscribe_event_err != nil {
				log.Println(unsubscribe_event_err)
			}

			email_vars := ConfirmationEmailTemplateVars{
				Code:     conf.Code,
				SiteName: opts.Config.Name,
				Paths:    opts.Config.Paths,
				Action:   "unsubscribe",
			}

			msg, err := message.NewMessageFromHTMLTemplate(email_t, email_vars)

			if err != nil {

				app_err := NewApplicationError(err, E_EMAIL_CREATE)
				vars.Error = app_err
				RenderTemplate(rsp, unsubscribe_t, vars)
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

			send_err := message.SendMessage(ctx, msg_opts, msg)

			send_event_params := url.Values{}
			send_event_params.Set("remote_addr", req.RemoteAddr)
			send_event_params.Set("confirmation_code", conf.Code)
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

			send_event_err := opts.EventLogs.AddEventLog(ctx, send_event)

			if send_event_err != nil {
				log.Println(send_event_err)
			}

			if send_err != nil {
				app_err := NewApplicationError(send_err, E_EMAIL_SEND)
				vars.Error = app_err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			RenderTemplate(rsp, success_t, vars)
			return

		default:
			http.Error(rsp, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	}

	h := http.HandlerFunc(fn)
	return h, nil
}
