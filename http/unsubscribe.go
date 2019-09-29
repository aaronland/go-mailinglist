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
	"github.com/aaronland/gomail"
	"html/template"
	"log"
	gohttp "net/http"
	"net/mail"
	"net/url"
	"time"
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

func UnsubscribeHandler(opts *UnsubscribeHandlerOptions) (gohttp.Handler, error) {

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

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		vars := UnsubscribeTemplateVars{
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
		}

		if !opts.FeatureFlags.Subscribe {
			vars.Error = errors.New("Disabled")
			RenderTemplate(rsp, unsubscribe_t, vars)
			return
		}

		switch req.Method {

		case "GET":

			RenderTemplate(rsp, unsubscribe_t, vars)
			return

		case "POST":

			subs_db := opts.Subscriptions
			conf_db := opts.Confirmations

			str_addr, err := sanitize.PostString(req, "address")

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			addr, err := mail.ParseAddress(str_addr)

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			sub, err := subs_db.GetSubscriptionWithAddress(addr.Address)

			if err != nil {

				if !database.IsNotExist(err) {
					vars.Error = err
					RenderTemplate(rsp, unsubscribe_t, vars)
					return
				}

				vars.Error = errors.New("Invalid subscription")
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			conf, err := confirmation.NewConfirmationForSubscription(sub, "unsubscribe")

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			err = conf_db.AddConfirmation(conf)

			if err != nil {

				vars.Error = err
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

			unsubscribe_event_err := opts.EventLogs.AddEventLog(unsubscribe_event)

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

				vars.Error = err
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

			send_err := message.SendMessage(msg, msg_opts)

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

			send_event_err := opts.EventLogs.AddEventLog(send_event)

			if send_event_err != nil {
				log.Println(send_event_err)
			}

			if send_err != nil {
				vars.Error = send_err
				RenderTemplate(rsp, unsubscribe_t, vars)
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
