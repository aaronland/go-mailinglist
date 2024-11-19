package main

import (
	"context"
	"flag"
	"log"

	aa_dynamodb "github.com/aaronland/go-aws-dynamodb"
	ml_dynamodb "github.com/aaronland/go-mailinglist/v2/dynamodb"
)

func main() {

	// var prefix string
	
	client_uri := flag.String("client-uri", "", "...")
	refresh := flag.Bool("refresh", false, "...")

	// flag.StringVar(&prefix, "prefix", "", "...")
	
	flag.Parse()

	ctx := context.Background()

	client, err := aa_dynamodb.NewClient(ctx, *client_uri)

	if err != nil {
		log.Fatalf("Failed to create new client, %w", err)
	}

	opts := &aa_dynamodb.CreateTablesOptions{
		Tables:  ml_dynamodb.DynamoDBTables,
		Refresh: *refresh,
		// Prefix: prefix,
	}

	err = aa_dynamodb.CreateTables(ctx, client, opts)

	if err != nil {
		log.Fatalf("Failed to create access tokens database, %v", err)
	}

}
