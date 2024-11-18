package main

import (
	"flag"
	"log"
	"time"
	"context"
	
	"github.com/aaronland/go-mailinglist/v2"
	"github.com/aaronland/go-mailinglist/v2/subscription"
)

func main() {

	subs_db_uri := flag.String("subscriptions-database-uri", "", "...")

	addr := flag.String("address", "", "...")
	confirmed := flag.Bool("confirmed", false, "...")

	flag.Parse()

	ctx := context.Background()
	
	db, err := database.NewSubscriptionsDatabase(ctx, *subs_db_uri)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	
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
