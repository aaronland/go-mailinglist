package http

import (
	"github.com/aaronland/go-mailinglist"
	gohttp "net/http"
)

func ConfirmHandler(subscriber_db mailinglist.SubscriptionDatabase) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {
		rsp.Header().Set("Content-Type", "text/plain")
		rsp.Write([]byte("PONG"))
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
