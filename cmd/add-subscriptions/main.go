package main

import (
	"context"
	"log"

	_ "github.com/aaronland/gocloud-docstore"
	_ "gocloud.dev/docstore/awsdynamodb"

	"github.com/aaronland/go-mailinglist/v2/app/subscriptions/add"
)

func main() {

	ctx := context.Background()
	err := add.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to add subscriptions, %v", err)
	}
}
