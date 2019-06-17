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

	subscribe := flag.Bool("subscribe", true, "...")
	unsubscribe := flag.Bool("unsubscribe", true, "...")

	path_subscribe := flag.String("path-unsubscribe", "/subscribe", "...")
	path_unsubscribe := flag.String("path-unsubscribe", "/unsubscribe", "...")
	path_ping := flag.String("path-ping", "/ping", "...")

	flag.Parse()

	db, err := mailinglist.NewSubscriberDatabaseFromDSN(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	mux := gohttp.NewServeMux()

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle(*path_ping, ping_handler)

	if *subscribe {

		subscribe_handler, err := http.SubscribeHandler(db)

		if err != nil {
			log.Fatal(err)
		}

		mux.Handle(*path_subscribe, subscribe_handler)
	}

	if *unsubscribe {

		unsubscribe_handler, err := http.UnsubscribeHandler(db)

		if err != nil {
			log.Fatal(err)
		}

		mux.Handle(*path_unsubscribe, unsubscribe_handler)
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
