package http

import (
	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist/confirmation"
	"github.com/aaronland/go-mailinglist/database"
	"html/template"
	gohttp "net/http"
	"net/mail"
)

type UnsubscribeTemplateVars struct {
	URL string
}

type UnsubscribeHandlerOptions struct {
	Subscriptions database.SubscriptionsDatabase
	Confirmations database.ConfirmationsDatabase
}

func UnsubscribeHandler(opts *UnsubscribeHandlerOptions) (gohttp.Handler, error) {

	unsubscribe_t := template.New("subscribe")

	unsubscribe_t, err := unsubscribe_t.Parse(`<html><head><title>Signup</title></head>
<body>
<form method="POST" action="{{ .URL }}">
<input type="text" name="address" id="address" placeholder="Enter your email address address" />
<button type="submit">Sign up</button>
</form>
</body></html>`)

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		switch req.Method {

		case "GET":

			vars := UnsubscribeTemplateVars{
				URL: req.URL.Path,
			}

			err := unsubscribe_t.Execute(rsp, vars)

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

				gohttp.Error(rsp, "NO SUB", gohttp.StatusBadRequest)
				return
			}

			conf, err := confirmation.NewConfirmationForSubscription(sub, "unsubscribe")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				return
			}

			err = conf_db.AddConfirmation(conf)

			if err != nil {

				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				return
			}

			// SEND CONFIRMATION CODE HERE...

			rsp.Write([]byte(conf.Code))
			return

		default:
			gohttp.Error(rsp, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}

	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
