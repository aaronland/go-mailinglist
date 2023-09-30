package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/subscription"
)

func main() {

	subs_uri := flag.String("subscriptions-uri", "", "...")

	addr := flag.String("address", "", "...")
	confirmed := flag.Bool("confirmed", false, "...")

	flag.Parse()

	ctx := context.Background()

	subs_db, err := database.NewSubscriptionsDatabase(ctx, *subs_uri)

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

	err = subs_db.AddSubscription(ctx, sub)

	if err != nil {
		log.Fatal(err)
	}

	if !sub.IsConfirmed() {
		// send confirmation code...
	}
}
