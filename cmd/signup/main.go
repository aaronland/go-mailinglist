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

	path_signup := flag.String("path-signup", "/signup", "...")
	path_ping := flag.String("path-ping", "/ping", "...")

	flag.Parse()

	db, err := mailinglist.NewSubscriberDatabaseFromDSN(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	signup_handler, err := http.SignupHandler(db)

	if err != nil {
		log.Fatal(err)
	}

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux := gohttp.NewServeMux()
	mux.Handle(*path_signup, signup_handler)
	mux.Handle(*path_ping, ping_handler)

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
