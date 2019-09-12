package http

import (
	"errors"
	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist/confirmation"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/message"
	"github.com/aaronland/gomail"
	"html/template"
	gohttp "net/http"
	"net/mail"
)

type UnsubscribeTemplateVars struct {
	URL   string
	Paths *PathOptions
	Error error
}

type UnsubscribeHandlerOptions struct {
	Templates     *template.Template
	Paths         *PathOptions
	Subscriptions database.SubscriptionsDatabase
	Confirmations database.ConfirmationsDatabase
	Sender        gomail.Sender
}

func UnsubscribeHandler(opts *UnsubscribeHandlerOptions) (gohttp.Handler, error) {

	unsubscribe_t, err := LoadTemplate(opts.Templates, "unsubscribe")

	if err != nil {
		return nil, err
	}

	confirm_t, err := LoadTemplate(opts.Templates, "unsubscribe_confirmation")

	if err != nil {
		return nil, err
	}

	email_t, err := LoadTemplate(opts.Templates, "email_confirmation")

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		vars := UnsubscribeTemplateVars{
			URL:   req.URL.Path,
			Paths: opts.Paths,
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

			email_vars := ConfirmationEmailTemplateVars{
				Code: conf.Code,
			}

			msg, err := message.NewMessageFromHTMLTemplate(email_t, email_vars)

			if err != nil {

				vars.Error = err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			from_addr, _ := mail.ParseAddress("fixme@localhost")
			to_addr, _ := mail.ParseAddress(sub.Address)

			msg_opts := &message.SendMessageOptions{
				Sender:  opts.Sender,
				Subject: "Your subscription...",
				From:    from_addr,
				To:      to_addr,
			}

			err = message.SendMessage(msg, msg_opts)

			if err != nil {

				vars.Error = err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			RenderTemplate(rsp, confirm_t, nil)
			return

		default:
			gohttp.Error(rsp, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}

	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
