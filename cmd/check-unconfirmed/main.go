package main

import (
	"flag"
	"github.com/aaronland/go-mailinglist"
	"log"
	_ "time"
)

func main() {

	dsn := flag.String("dsn", "", "...")

	flag.Parse()

	db, err := mailinglist.NewSubscriptionDatabaseFromDSN(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	cb := func(sub *mailinglist.Subscriber) error {

		// where/how to check status here?
		return nil
	}

	db.UnconfirmedSubscriptions(cb)
}
