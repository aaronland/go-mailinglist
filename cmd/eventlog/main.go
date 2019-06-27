package main

import (
	"flag"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/eventlog"
	"log"
	"os"
)

func main() {

	subs_dsn := flag.String("subscriptions-dsn", "", "...")
	logs_dsn := flag.String("eventlogs-dsn", "", "...")
	message := flag.String("message", "", "...")

	addr := flag.String("address", "", "...")

	flag.Parse()

	if *message == "" {
		log.Fatal("Invalid -message parameter")
	}

	subs_db, err := mailinglist.NewSubscriptionsDatabaseFromDSN(*subs_dsn)

	if err != nil {
		log.Fatal(err)
	}

	logs_db, err := mailinglist.NewEventLogsDatabaseFromDSN(*logs_dsn)

	if err != nil {
		log.Fatal(err)
	}

	sub, err := subs_db.GetSubscriptionWithAddress(*addr)

	if err != nil {
		log.Fatal(err)
	}

	event_log, err := eventlog.NewEventLogWithSubscription(sub, eventlog.EVENTLOG_CUSTOM_EVENT, *message)

	if err != nil {
		log.Fatal(err)
	}

	err = logs_db.AddEventLog(event_log)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
