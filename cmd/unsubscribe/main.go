package main

import (
	"flag"
	"log"
	"context"

	"github.com/aaronland/go-mailinglist/v2/database"		
)

func main() {

	subs_db_uri := flag.String("subscriptions-database-uri", "", "...")
	addr := flag.String("address", "", "...")

	flag.Parse()

	db, err := database.NewSubscriptionsDatabase(ctx, *subs_db_uri)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	
	sub, err := db.GetSubscriptionWithAddress(ctx, *addr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.RemoveSubscription(ctx, sub)

	if err != nil {
		log.Fatal(err)
	}
}
