package http

import (
	"github.com/aaronland/go-mailinglist/database"
	"html/template"
	gohttp "net/http"
)

type SubscribeVars struct {
	URL string
}

func SubscribeHandler(subs_db database.SubscriptionsDatabase) (gohttp.Handler, error) {

	subscribe_t := template.New("subscribe")
	
	subscribe_t, err := subscribe_t.Parse(`<html><head><title>Signup</title></head>
<body>
<form method="POST" action="{{ .URL }}">
<input type="text" name="email" id="email" placeholder="Enter your email address address" />
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
			// pass
		default:
			gohttp.Error(rsp, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
