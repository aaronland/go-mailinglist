package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var DynamoDBSubscriptionsTable = &dynamodb.CreateTableInput{
	AttributeDefinitions: []types.AttributeDefinition{
		{
			AttributeName: aws.String("Address"),
			AttributeType: "S",
		},
		{
			AttributeName: aws.String("Status"),
			AttributeType: "N",
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
			IndexName: aws.String("status"),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("Status"),
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
