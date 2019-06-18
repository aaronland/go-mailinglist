package main

import (
	"flag"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/subscription"
	"log"
	"time"
)

func main() {

	dsn := flag.String("dsn", "", "...")

	addr := flag.String("address", "", "...")
	confirmed := flag.Bool("confirmed", false, "...")

	flag.Parse()

	db, err := mailinglist.NewSubscriptionsDatabaseFromDSN(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	sub, err := subscription.NewSubscription(*addr)

	if err != nil {
		log.Fatal(err)
	}

	if *confirmed {
		now := time.Now()
		sub.Confirmed = now.Unix()
	}

	err = db.AddSubscription(sub)

	if err != nil {
		log.Fatal(err)
	}

	if !sub.IsConfirmed() {
		// send confirmation code...
	}
}
