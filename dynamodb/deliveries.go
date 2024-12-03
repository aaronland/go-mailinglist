package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var DynamoDBDeliveriesTable = &dynamodb.CreateTableInput{
	AttributeDefinitions: []types.AttributeDefinition{
		{
			AttributeName: aws.String("Address"),
			AttributeType: "S",
		},
		{
			AttributeName: aws.String("MessageId"),
			AttributeType: "S",
		},
		{
			AttributeName: aws.String("Delivered"),
			AttributeType: "N",
		},
	},
	KeySchema: []types.KeySchemaElement{
		{
			AttributeName: aws.String("Address"),
			KeyType:       "HASH",
		},
		{
			AttributeName: aws.String("Delivered"),
			KeyType:       "RANGE",
		},
	},
	GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
		{
			IndexName: aws.String("address_message"),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("Address"),
					KeyType:       "HASH",
				},
				{
					AttributeName: aws.String("MessageId"),
					KeyType:       "RANGE",
				},
			},
			Projection: &types.Projection{
				// maybe just address...?
				ProjectionType: "ALL",
			},
		},
	},
	BillingMode: types.BillingModePayPerRequest,
}
