package main

import (
	"context"
	"log"

	"github.com/aaronland/go-mailinglist/v2/app/subscriptions/list"
)

func main() {

	ctx := context.Background()
	err := list.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to list subscriptions, %v", err)
	}
}
