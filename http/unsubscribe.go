package http

import (
	"github.com/aaronland/go-mailinglist"
	gohttp "net/http"
)

func UnsubscribeHandler(subscriber_db mailinglist.SubscriberDatabase) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {
		rsp.Header().Set("Content-Type", "text/plain")
		rsp.Write([]byte("PONG"))
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
