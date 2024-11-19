package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var DynamoDBSubscriptionsTable = &dynamodb.CreateTableInput{
	AttributeDefinitions: []types.AttributeDefinition{
		{
			AttributeName: aws.String("address"),
			AttributeType: "S",
		},
		{
			AttributeName: aws.String("status"),
			AttributeType: "N",
		},
	},
	KeySchema: []types.KeySchemaElement{
		{
			AttributeName: aws.String("address"),
			KeyType:       "HASH",
		},
	},
	GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
		{
			IndexName: aws.String("status"),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("status"),
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
