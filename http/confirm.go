package http

import (
	"errors"
	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/database"
	"html/template"
	_ "log"
	gohttp "net/http"
	"time"
)

type ConfirmHandlerOptions struct {
	Config        *mailinglist.MailingListConfig
	Templates     *template.Template
	Subscriptions database.SubscriptionsDatabase
	Confirmations database.ConfirmationsDatabase
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

				RenderTemplate(rsp, success_t, vars)
				return

			case "unsubscribe":

				err = subs_db.RemoveSubscription(sub)

				if err != nil {
					vars.Error = err
					RenderTemplate(rsp, action_t, vars)
					return
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
