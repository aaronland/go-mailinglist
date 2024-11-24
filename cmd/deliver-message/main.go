package main

import (
	"context"
	"log"

	_ "github.com/aaronland/gocloud-blob/s3"
	_ "github.com/aaronland/gocloud-docstore"
	_ "github.com/aaronland/gomail-sender-ses/v2"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/s3blob"
	_ "gocloud.dev/docstore/awsdynamodb"

	"github.com/aaronland/go-mailinglist/v2/app/message/deliver"
)

func main() {

	ctx := context.Background()
	err := deliver.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to deliver message, %v", err)
	}
}
