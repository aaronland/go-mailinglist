package main

import (
	"flag"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/subscription"
	"log"
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

	sub.Confirmed = *confirmed

	err = db.AddSubscription(sub)

	if err != nil {
		log.Fatal(err)
	}

	if sub.Confirmed == false {
		// send confirmation code...
	}
}
