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
	},
	KeySchema: []types.KeySchemaElement{
		{
			AttributeName: aws.String("Address"),
			KeyType:       "HASH",
		},
	},
	GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
		{
			IndexName: aws.String("message_id"),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("MessageId"),
					KeyType:       "HASH",
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
