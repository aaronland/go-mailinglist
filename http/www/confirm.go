package www

// CSRF crumbs are handled by go-http-crumb middleware
// Bootstrap stuff is handled by go-http-bootstrap middleware
// see cmd/subscriptiond/main.go for details

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist/v2"
	"github.com/aaronland/go-mailinglist/v2/database"
	"github.com/aaronland/go-mailinglist/v2/eventlog"
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

func ConfirmHandler(opts *ConfirmHandlerOptions) (http.Handler, error) {

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

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		subs_db := opts.Subscriptions
		conf_db := opts.Confirmations

		vars := ConfirmTemplateVars{
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
		}

		if !opts.Config.FeatureFlags.Subscribe {

			app_err := NewApplicationError(nil, E_DISABLED_SUBSCRIBE)
			vars.Error = app_err

			RenderTemplate(rsp, confirm_t, vars)
			return
		}

		switch req.Method {

		case "GET":

			code, err := sanitize.GetString(req, "code")

			if err != nil {

				app_err := NewApplicationError(err, E_INPUT_PARSE, "code")
				vars.Error = app_err

				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			if code == "" {

				app_err := NewApplicationError(nil, E_INPUT_PARSE, "code")
				vars.Error = app_err

				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			vars.Code = code

			conf, err := conf_db.GetConfirmationWithCode(ctx, code)

			if err != nil {

				app_err := NewApplicationError(err, E_CONFIRMATION_RETRIEVE)
				vars.Error = app_err
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			if conf.IsExpired() {

				app_err := NewApplicationError(nil, E_CONFIRMATION_EXPIRED)
				vars.Error = app_err

				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			_, err = subs_db.GetSubscriptionWithAddress(ctx, conf.Address)

			if err != nil {
				app_err := NewApplicationError(err, E_SUBSCRIPTION_RETRIEVE)
				vars.Error = app_err
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			/*
				if !sub.IsEnabled(){
					vars.Error = errors.New("Disabled")
					RenderTemplate(rsp, confirm_t, vars)
					return
				}
			*/

			vars.Action = conf.Action
			RenderTemplate(rsp, action_t, vars)

			return

		case "POST":

			code, err := sanitize.PostString(req, "code")

			if err != nil {
				app_err := NewApplicationError(err, E_INPUT_PARSE, "code")
				vars.Error = app_err
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			vars.Code = code

			confirmed, err := sanitize.PostString(req, "confirm")

			if err != nil {
				app_err := NewApplicationError(err, E_INPUT_PARSE, "confirm")
				vars.Error = app_err
				RenderTemplate(rsp, action_t, vars)
				return
			}

			conf, err := conf_db.GetConfirmationWithCode(ctx, code)

			if err != nil {
				app_err := NewApplicationError(err, E_CONFIRMATION_RETRIEVE)
				vars.Error = app_err
				RenderTemplate(rsp, action_t, vars)
				return
			}

			if conf.IsExpired() {
				app_err := NewApplicationError(nil, E_CONFIRMATION_EXPIRED)
				vars.Error = app_err
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

			vars.Action = conf.Action

			if confirmed == "" {
				RenderTemplate(rsp, action_t, vars)
				return
			}

			sub, err := subs_db.GetSubscriptionWithAddress(ctx, conf.Address)

			if err != nil {
				app_err := NewApplicationError(err, E_SUBSCRIPTION_RETRIEVE)
				vars.Error = app_err
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

				err = sub.Confirm()

				if err != nil {
					app_err := NewApplicationError(err, E_CONFIRMATION_CONFIRM)
					vars.Error = app_err
					RenderTemplate(rsp, action_t, vars)
					return
				}

				err = subs_db.UpdateSubscription(ctx, sub)

				if err != nil {
					app_err := NewApplicationError(err, E_SUBSCRIPTION_UPDATE)
					vars.Error = app_err
					RenderTemplate(rsp, action_t, vars)
					return
				}

				confirm_event_err := opts.EventLogs.AddEventLog(ctx, confirm_event)

				if confirm_event_err != nil {
					log.Println(confirm_event_err)
				}

				RenderTemplate(rsp, success_t, vars)
				return

			case "unsubscribe":

				if !sub.IsEnabled() {
					app_err := NewApplicationError(nil, E_SUBSCRIPTION_DISABLED)
					vars.Error = app_err
					RenderTemplate(rsp, action_t, vars)
					return
				}

				err = subs_db.RemoveSubscription(ctx, sub)

				if err != nil {
					app_err := NewApplicationError(err, E_SUBSCRIPTION_REMOVE)
					vars.Error = app_err
					RenderTemplate(rsp, action_t, vars)
					return
				}

				confirm_event_err := opts.EventLogs.AddEventLog(ctx, confirm_event)

				if confirm_event_err != nil {
					log.Println(confirm_event_err)
				}

				RenderTemplate(rsp, success_t, vars)
				return

			default:

				app_err := NewApplicationError(nil, E_CONFIRMATION_INVALID)
				vars.Error = app_err
				RenderTemplate(rsp, confirm_t, vars)
				return
			}

		default:
			http.Error(rsp, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
