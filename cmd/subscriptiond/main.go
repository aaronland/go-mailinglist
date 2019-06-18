package main

import (
	"flag"
	"github.com/aaronland/go-http-crumb"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/http"
	"github.com/aaronland/go-mailinglist/server"
	"log"
	gohttp "net/http"
)

func main() {

	subs_dsn := flag.String("subscriptions-dsn", "", "...")
	conf_dsn := flag.String("confirmations-dsn", "", "...")
	crumb_dsn := flag.String("crumb-dsn", "", "...")

	protocol := flag.String("protocol", "http", "...")
	host := flag.String("host", "localhost", "...")
	port := flag.Int("port", 8080, "...")

	subscribe_handler := flag.Bool("subscribe-handler", true, "...")
	unsubscribe_handler := flag.Bool("unsubscribe-handler", true, "...")
	confirm_handler := flag.Bool("confirm-handler", true, "...")

	path_subscribe := flag.String("path-subscribe", "/subscribe", "...")
	path_unsubscribe := flag.String("path-unsubscribe", "/unsubscribe", "...")
	path_confirm := flag.String("path-confirm", "/confirm", "...")

	path_ping := flag.String("path-ping", "/ping", "...")

	flag.Parse()

	subs_db, err := mailinglist.NewSubscriptionsDatabaseFromDSN(*subs_dsn)

	if err != nil {
		log.Fatal(err)
	}

	conf_db, err := mailinglist.NewConfirmationsDatabaseFromDSN(*conf_dsn)

	if err != nil {
		log.Fatal(err)
	}

	mux := gohttp.NewServeMux()

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	crumb_cfg, err := crumb.NewCrumbConfigFromDSN(*crumb_dsn)

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle(*path_ping, ping_handler)

	if *subscribe_handler {

		h, err := http.SubscribeHandler(subs_db, conf_db)

		if err != nil {
			log.Fatal(err)
		}

		cr := crumb.EnsureCrumbHandler(crumb_cfg, h)

		mux.Handle(*path_subscribe, cr)
	}

	if *unsubscribe_handler {

		h, err := http.UnsubscribeHandler(subs_db)

		if err != nil {
			log.Fatal(err)
		}

		cr := crumb.EnsureCrumbHandler(crumb_cfg, h)

		mux.Handle(*path_unsubscribe, cr)
	}

	if *confirm_handler {

		h, err := http.ConfirmHandler(subs_db)

		if err != nil {
			log.Fatal(err)
		}

		cr := crumb.EnsureCrumbHandler(crumb_cfg, h)

		mux.Handle(*path_confirm, cr)
	}

	s, err := server.NewServer(*protocol, *host, *port)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s\n", s.Address())

	err = s.ListenAndServe(mux)

	if err != nil {
		log.Fatal(err)
	}
}
