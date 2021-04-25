package main

import (
	"context"
	"flag"
	"github.com/aaronland/go-mailinglist/database"
	"log"
)

func main() {

	subs_uri := flag.String("subscriptions-uri", "", "...")
	addr := flag.String("address", "", "...")

	flag.Parse()

	ctx := context.Background()

	subs_db, err := database.NewSubscriptionsDatabase(ctx, *subs_uri)

	if err != nil {
		log.Fatal(err)
	}

	sub, err := subs_db.GetSubscriptionWithAddress(*addr)

	if err != nil {
		log.Fatal(err)
	}

	err = subs_db.RemoveSubscription(sub)

	if err != nil {
		log.Fatal(err)
	}
}
