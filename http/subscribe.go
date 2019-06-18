package http

import (
	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/subscription"
	"html/template"
	_ "log"
	gohttp "net/http"
	"net/mail"
)

type SubscribeVars struct {
	URL string
}

func SubscribeHandler(subs_db database.SubscriptionsDatabase) (gohttp.Handler, error) {

	subscribe_t := template.New("subscribe")

	subscribe_t, err := subscribe_t.Parse(`<html><head><title>Signup</title></head>
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

			vars := SubscribeVars{
				URL: req.URL.Path,
			}

			err := subscribe_t.Execute(rsp, vars)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			}

			return

		case "POST":

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

			// check is not exist here...

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
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

			rsp.Write([]byte(sub.Address))
			return

		default:
			gohttp.Error(rsp, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
