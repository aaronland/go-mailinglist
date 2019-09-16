package http

import (
	"errors"
	"fmt"
	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/confirmation"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/message"
	"github.com/aaronland/gomail"
	"html/template"
	_ "log"
	gohttp "net/http"
	"net/mail"
)

type UnsubscribeTemplateVars struct {
	URL      string
	SiteName string
	Paths    *mailinglist.PathConfig
	Error    error
}

type UnsubscribeHandlerOptions struct {
	Config        *mailinglist.MailingListConfig
	Templates     *template.Template
	Subscriptions database.SubscriptionsDatabase
	Confirmations database.ConfirmationsDatabase
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
			URL:      req.URL.Path,
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
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
				Code:     conf.Code,
				URL:      req.URL.Path,
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

			err = message.SendMessage(msg, msg_opts)
			
			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, unsubscribe_t, vars)
				return
			}

			RenderTemplate(rsp, success_t, nil)
			return

		default:
			gohttp.Error(rsp, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}

	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
