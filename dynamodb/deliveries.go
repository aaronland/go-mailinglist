package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var DynamoDBDeliveriesTable = &dynamodb.CreateTableInput{
	AttributeDefinitions: []types.AttributeDefinition{
		{
			AttributeName: aws.String("address"),
			AttributeType: "S",
		},
		{
			AttributeName: aws.String("message_id"),
			AttributeType: "S",
		},
	},
	KeySchema: []types.KeySchemaElement{
		{
			AttributeName: aws.String("address"),
			KeyType:       "HASH",
		},
		{
			AttributeName: aws.String("message_id"),
			KeyType:       "RANGE",
		},
	},
	GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
		{
			IndexName: aws.String("status"),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("message_id"),
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
