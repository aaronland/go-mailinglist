package main

import (
	"flag"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/http"
	"github.com/aaronland/go-mailinglist/server"
	"log"
	gohttp "net/http"
)

func main() {

	dsn := flag.String("dsn", "", "...")

	protocol := flag.String("protocol", "http", "...")
	host := flag.String("host", "localhost", "...")
	port := flag.Int("port", 8080, "...")

	subscribe_handler := flag.Bool("subscribe", true, "...")
	unsubscribe_handler := flag.Bool("unsubscribe", true, "...")
	confirm_handler := flag.Bool("confirm-handler", true, "...")

	path_subscribe := flag.String("path-unsubscribe", "/subscribe", "...")
	path_unsubscribe := flag.String("path-unsubscribe", "/unsubscribe", "...")
	path_confirm := flag.String("path-confirm", "/confirm", "...")

	path_ping := flag.String("path-ping", "/ping", "...")

	flag.Parse()

	db, err := mailinglist.NewSubscriptionsDatabaseFromDSN(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	mux := gohttp.NewServeMux()

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle(*path_ping, ping_handler)

	if *subscribe_handler {

		h, err := http.SubscribeHandler(db)

		if err != nil {
			log.Fatal(err)
		}

		mux.Handle(*path_subscribe, h)
	}

	if *unsubscribe_handler {

		h, err := http.UnsubscribeHandler(db)

		if err != nil {
			log.Fatal(err)
		}

		mux.Handle(*path_unsubscribe, h)
	}

	if *confirm_handler {

		h, err := http.ConfirmHandler(db)

		if err != nil {
			log.Fatal(err)
		}

		mux.Handle(*path_confirm, h)
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
