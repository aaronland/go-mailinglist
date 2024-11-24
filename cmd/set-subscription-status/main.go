package main

import (
	"context"
	"log"

	_ "github.com/aaronland/gocloud-docstore"
	_ "gocloud.dev/docstore/awsdynamodb"

	"github.com/aaronland/go-mailinglist/v2/app/subscriptions/status"
)

func main() {

	ctx := context.Background()
	err := status.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to set subscription status, %v", err)
	}
}
