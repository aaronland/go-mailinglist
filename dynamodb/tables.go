package dynamodb

import (
	aws_dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DynamoDBTables = map[string]*aws_dynamodb.CreateTableInput{
	"subscriptions": DynamoDBSubscriptionsTable,
	"confirmations": DynamoDBConfirmationsTable,
	"deliveries":    DynamoDBDeliveriesTable,
	// "invitations": InvitationsTable,
	"eventlogs": DynamoDBEventLogsTable,
}
