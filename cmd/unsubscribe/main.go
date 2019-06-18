package main

import (
	"flag"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/database"
	"log"
)

func main() {

	dsn := flag.String("dsn", "", "...")
	addr := flag.String("address", "", "...")

	flag.Parse()

	db, err := database.NewSubscriptionDatabaseFromDSN(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	sub, err := db.GetSubscriptionWithAddress(*addr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.AddSubscription(sub)

	if err != nil {
		log.Fatal(err)
	}
}
