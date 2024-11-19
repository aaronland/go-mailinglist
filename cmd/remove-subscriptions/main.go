package main

import (
	"context"
	"log"

	"github.com/aaronland/go-mailinglist/v2/app/subscriptions/remove"
)

func main() {

	ctx := context.Background()
	err := remove.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to remove subscriptions, %v", err)
	}
}
