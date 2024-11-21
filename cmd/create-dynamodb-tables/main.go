package main

import (
	"context"
	"flag"
	"log"

	aa_dynamodb "github.com/aaronland/go-aws-dynamodb"
	"github.com/aaronland/go-mailinglist/v2/dynamodb"
)

func main() {

	var client_uri string
	var refresh bool
	var prefix string

	flag.StringVar(&client_uri, "client-uri", "aws://?region=localhost&credentials=anon:&local=true", "...")
	flag.BoolVar(&refresh, "refresh", false, "...")
	flag.StringVar(&prefix, "prefix", "", "Optional string to prepend to all table names")

	flag.Parse()

	ctx := context.Background()

	client, err := aa_dynamodb.NewClient(ctx, client_uri)

	if err != nil {
		log.Fatalf("Failed to create client, %v", err)
	}

	table_opts := &aa_dynamodb.CreateTablesOptions{
		Tables:  dynamodb.DynamoDBTables,
		Refresh: refresh,
		Prefix:  prefix,
	}

	err = aa_dynamodb.CreateTables(ctx, client, table_opts)

	if err != nil {
		log.Fatalf("Failed to create tables, %v", err)
	}
}
