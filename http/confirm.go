package http

import (
	"github.com/aaronland/go-mailinglist/database"
	gohttp "net/http"
)

type ConfirmHandlerOptions struct {
	Subscriptions database.SubscriptionsDatabase
	Confirmations database.ConfirmationsDatabase
}

func ConfirmHandler(opts *ConfirmHandlerOptions) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {
		rsp.Header().Set("Content-Type", "text/plain")
		rsp.Write([]byte("PONG"))
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
