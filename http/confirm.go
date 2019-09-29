package http

// CSRF crumbs are handled by go-http-crumb middleware
// Bootstrap stuff is handled by go-http-bootstrap middleware
// see cmd/subscriptiond/main.go for details

import (
	"errors"
	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/eventlog"
	"html/template"
	"log"
	gohttp "net/http"
	"net/url"
	"time"
)

type ConfirmHandlerOptions struct {
	Config        *mailinglist.MailingListConfig
	Templates     *template.Template
	Subscriptions database.SubscriptionsDatabase
	Confirmations database.ConfirmationsDatabase
	EventLogs     database.EventLogsDatabase
}

type ConfirmTemplateVars struct {
	SiteName string
	Paths    *mailinglist.PathConfig
	Code     string
	Action   string
	Error    error
}

func ConfirmHandler(opts *ConfirmHandlerOptions) (gohttp.Handler, error) {

	confirm_t, err := LoadTemplate(opts.Templates, "confirm")

	if err != nil {
		return nil, err
	}

	action_t, err := LoadTemplate(opts.Templates, "confirm_action")

	if err != nil {
		return nil, err
	}

	success_t, err := LoadTemplate(opts.Templates, "confirm_success")

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		subs_db := opts.Subscriptions
		conf_db := opts.Confirmations

		vars := ConfirmTemplateVars{
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
		}

		if !opts.Config.FeatureFlags.Subscribe {
			vars.Error = errors.New("Disabled")
			RenderTemplate(rsp, confirm_t, vars)
			return
		}

		switch req.Method {

		case "GET":

			code, err := sanitize.GetString(req, "code")

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			if code == "" {
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			vars.Code = code

			conf, err := conf_db.GetConfirmationWithCode(code)

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			if conf.IsExpired() {
				vars.Error = errors.New("Confirmation code is expired.")
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			_, err = subs_db.GetSubscriptionWithAddress(conf.Address)

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			vars.Action = conf.Action
			RenderTemplate(rsp, action_t, vars)

			return

		case "POST":

			code, err := sanitize.PostString(req, "code")

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			vars.Code = code

			confirmed, err := sanitize.PostString(req, "confirm")

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, action_t, vars)
				return
			}

			conf, err := conf_db.GetConfirmationWithCode(code)

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, action_t, vars)
				return
			}

			if conf.IsExpired() {
				vars.Error = errors.New("Expired")
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			vars.Action = conf.Action

			if confirmed == "" {
				RenderTemplate(rsp, action_t, vars)
				return
			}

			sub, err := subs_db.GetSubscriptionWithAddress(conf.Address)

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, action_t, vars)
				return
			}

			confirm_event_params := url.Values{}
			confirm_event_params.Set("remote_addr", req.RemoteAddr)
			confirm_event_params.Set("action", conf.Action)
			confirm_event_params.Set("confirmation_code", conf.Code)

			confirm_event_message := confirm_event_params.Encode()

			confirm_event := &eventlog.EventLog{
				Address: sub.Address,
				Created: time.Now().Unix(),
				Event:   eventlog.EVENTLOG_CONFIRM_EVENT,
				Message: confirm_event_message,
			}

			switch conf.Action {
			case "subscribe":

				now := time.Now()
				sub.Confirmed = now.Unix()

				err = subs_db.UpdateSubscription(sub)

				if err != nil {
					vars.Error = err
					RenderTemplate(rsp, action_t, vars)
					return
				}

				confirm_event_err := opts.EventLogs.AddEventLog(confirm_event)

				if confirm_event_err != nil {
					log.Println(confirm_event_err)
				}

				RenderTemplate(rsp, success_t, vars)
				return

			case "unsubscribe":

				err = subs_db.RemoveSubscription(sub)

				if err != nil {
					vars.Error = err
					RenderTemplate(rsp, action_t, vars)
					return
				}

				confirm_event_err := opts.EventLogs.AddEventLog(confirm_event)

				if confirm_event_err != nil {
					log.Println(confirm_event_err)
				}

				RenderTemplate(rsp, success_t, vars)
				return

			default:

				vars.Error = errors.New("Invalid action")
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

		default:
			gohttp.Error(rsp, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
