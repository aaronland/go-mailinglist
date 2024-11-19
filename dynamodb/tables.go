package dynamodb

import (
	aws_dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DynamoDBTables = map[string]*aws_dynamodb.CreateTableInput{
	/*
		"subscriptions":   SubscriptionsTable,
		"confirmations":   ConfirmationsTable,
		"deliveries": DeliveriesTable,
		"invitations": InvitationsTable,
		"eventlogs": EventLogsTable,
	*/
}
