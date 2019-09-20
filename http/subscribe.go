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
	"github.com/aaronland/go-mailinglist/message"
	"github.com/aaronland/go-mailinglist/subscription"
	"github.com/aaronland/gomail"
	"html/template"
	_ "log"
	gohttp "net/http"
	"net/mail"
)

type SubscribeTemplateVars struct {
	SiteName string
	Paths    *mailinglist.PathConfig
	Error    error
}

type SubscribeHandlerOptions struct {
	Config        *mailinglist.MailingListConfig
	Templates     *template.Template
	Subscriptions database.SubscriptionsDatabase
	Confirmations database.ConfirmationsDatabase
	Sender        gomail.Sender
}

func SubscribeHandler(opts *SubscribeHandlerOptions) (gohttp.Handler, error) {

	subscribe_t, err := LoadTemplate(opts.Templates, "subscribe")

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

		vars := SubscribeTemplateVars{
			SiteName: opts.Config.Name,
			Paths:    opts.Config.Paths,
		}

		switch req.Method {

		case "GET":
			RenderTemplate(rsp, subscribe_t, vars)
			return

		case "POST":

			subs_db := opts.Subscriptions
			conf_db := opts.Confirmations

			str_addr, err := sanitize.PostString(req, "address")

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, subscribe_t, vars)
				return
			}

			if str_addr == "" {
				vars.Error = errors.New("Empty address")
				RenderTemplate(rsp, subscribe_t, vars)
				return
			}

			addr, err := mail.ParseAddress(str_addr)

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, subscribe_t, vars)
				return
			}

			sub, err := subs_db.GetSubscriptionWithAddress(addr.Address)

			if err != nil {

				if !database.IsNotExist(err) {
					vars.Error = err
					RenderTemplate(rsp, subscribe_t, vars)
					return
				}
			}

			if sub != nil {
				rsp.Write([]byte("EXISTS"))
				return
			}

			sub, err = subscription.NewSubscription(addr.Address)

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, subscribe_t, vars)
				return
			}

			conf, err := confirmation.NewConfirmationForSubscription(sub, "subscribe")

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, subscribe_t, vars)
				return
			}

			err = subs_db.AddSubscription(sub)

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, subscribe_t, vars)
				return
			}

			err = conf_db.AddConfirmation(conf)

			if err != nil {

				go subs_db.RemoveSubscription(sub)

				vars.Error = err
				RenderTemplate(rsp, subscribe_t, vars)
				return
			}

			email_vars := ConfirmationEmailTemplateVars{
				Code:     conf.Code,
				SiteRoot: "fix me",
				SiteName: opts.Config.Name,
				Paths:    opts.Config.Paths,
				Action:   "subscribe",
			}

			msg, err := message.NewMessageFromHTMLTemplate(email_t, email_vars)

			if err != nil {
				vars.Error = err
				RenderTemplate(rsp, subscribe_t, vars)
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
				RenderTemplate(rsp, subscribe_t, vars)
				return
			}

			if err != nil {

				vars.Error = err
				RenderTemplate(rsp, subscribe_t, vars)
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
