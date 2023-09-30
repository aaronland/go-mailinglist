package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/eventlog"	
)

func main() {

	subs_uri := flag.String("subscriptions-uri", "", "...")
	logs_uri := flag.String("eventlogs-uri", "", "...")
	message := flag.String("message", "", "...")

	addr := flag.String("address", "", "...")

	flag.Parse()

	if *message == "" {
		log.Fatal("Invalid -message parameter")
	}

	ctx := context.Background()

	subs_db, err := database.NewSubscriptionsDatabase(ctx, *subs_uri)

	if err != nil {
		log.Fatal(err)
	}

	logs_db, err := database.NewEventLogsDatabase(ctx, *logs_uri)

	if err != nil {
		log.Fatal(err)
	}

	sub, err := subs_db.GetSubscriptionWithAddress(ctx, *addr)

	if err != nil {
		log.Fatal(err)
	}

	event_log, err := eventlog.NewEventLogWithSubscription(sub, eventlog.EVENTLOG_CUSTOM_EVENT, *message)

	if err != nil {
		log.Fatal(err)
	}

	err = logs_db.AddEventLog(ctx, event_log)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
