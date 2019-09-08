package http

import (
	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist/database"
	"html/template"
	gohttp "net/http"
	"time"
)

type ConfirmHandlerOptions struct {
	Templates     *template.Template
	Subscriptions database.SubscriptionsDatabase
	Confirmations database.ConfirmationsDatabase
}

type ConfirmTemplateVars struct {
	Code string
}

func ConfirmHandler(opts *ConfirmHandlerOptions) (gohttp.Handler, error) {

	confirm_t, err := LoadTemplate(opts.Templates, "confirm")

	if err != nil {
		return nil, err
	}

	update_t, err := LoadTemplate(opts.Templates, "confirm_update")

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		subs_db := opts.Subscriptions
		conf_db := opts.Confirmations

		switch req.Method {

		case "GET":

			code, err := sanitize.GetString(req, "code")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			conf, err := conf_db.GetConfirmationWithCode(code)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			if conf.IsExpired() {
				gohttp.Error(rsp, "EXPIRED", gohttp.StatusBadRequest)
				return
			}

			_, err = subs_db.GetSubscriptionWithAddress(conf.Address)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			confirm_vars := ConfirmTemplateVars{
				Code: code,
			}

			err = confirm_t.Execute(rsp, confirm_vars)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			}

			return

		case "POST":

			code, err := sanitize.PostString(req, "code")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			_, err = sanitize.PostString(req, "confirm")

			if err != nil {
				// FIX ME : REDIRECT TO GET...
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			conf, err := conf_db.GetConfirmationWithCode(code)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			if conf.IsExpired() {
				gohttp.Error(rsp, "EXPIRED", gohttp.StatusBadRequest)
				return
			}

			sub, err := subs_db.GetSubscriptionWithAddress(conf.Address)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			switch conf.Action {
			case "subscribe":

				now := time.Now()
				sub.Confirmed = now.Unix()

				err = subs_db.UpdateSubscription(sub)

				if err != nil {
					gohttp.Error(rsp, "Invalid action", gohttp.StatusInternalServerError)
					return
				}

				err = update_t.Execute(rsp, nil)

				if err != nil {
					gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				}

				return

			case "unsubscribe":

				err = subs_db.RemoveSubscription(sub)

				if err != nil {
					gohttp.Error(rsp, "Invalid action", gohttp.StatusInternalServerError)
					return
				}

				err = update_t.Execute(rsp, nil)

				if err != nil {
					gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				}

				return

			default:
				gohttp.Error(rsp, "Invalid action", gohttp.StatusInternalServerError)
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
