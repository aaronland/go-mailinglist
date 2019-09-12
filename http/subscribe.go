package http

import (
	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist/confirmation"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/message"
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/aaronland/gomail"
	"html/template"
	_ "log"
	gohttp "net/http"
	"net/mail"
)

type SubscribeTemplateVars struct {
	URL string
	Paths *PathOptions
	Error error
}

type ConfirmationEmailTemplateVars struct {
	Code string
}

type SubscribeHandlerOptions struct {
	Templates     *template.Template
	Paths         *PathOptions
	Subscriptions database.SubscriptionsDatabase
	Confirmations database.ConfirmationsDatabase
	Sender        gomail.Sender
}

func SubscribeHandler(opts *SubscribeHandlerOptions) (gohttp.Handler, error) {

	subscribe_t, err := LoadTemplate(opts.Templates, "subscribe")

	if err != nil {
		return nil, err
	}

	confirm_t, err := LoadTemplate(opts.Templates, "subscribe_confirmation")

	if err != nil {
		return nil, err
	}

	email_t, err := LoadTemplate(opts.Templates, "email_confirmation")

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		switch req.Method {

		case "GET":

			vars := SubscribeTemplateVars{
				URL: req.URL.Path,
				Paths: opts.Paths,
			}

			err := subscribe_t.Execute(rsp, vars)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			}

			return

		case "POST":

			subs_db := opts.Subscriptions
			conf_db := opts.Confirmations

			str_addr, err := sanitize.PostString(req, "address")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			addr, err := mail.ParseAddress(str_addr)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			sub, err := subs_db.GetSubscriptionWithAddress(addr.Address)

			if err != nil {

				if !database.IsNotExist(err) {
					gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
					return
				}
			}

			if sub != nil {
				rsp.Write([]byte("EXISTS"))
				return
			}

			sub, err = subscription.NewSubscription(addr.Address)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				return
			}

			conf, err := confirmation.NewConfirmationForSubscription(sub, "subscribe")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				return
			}

			err = subs_db.AddSubscription(sub)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				return
			}

			err = conf_db.AddConfirmation(conf)

			if err != nil {

				go subs_db.RemoveSubscription(sub)

				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				return
			}

			email_vars := ConfirmationEmailTemplateVars{
				Code: conf.Code,
			}

			msg, err := message.NewMessageFromHTMLTemplate(email_t, email_vars)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
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
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				return
			}

			err = confirm_t.Execute(rsp, nil)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			}

			return

		default:
			gohttp.Error(rsp, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
