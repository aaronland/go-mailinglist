package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var DynamoDBEventLogsTable = &dynamodb.CreateTableInput{
	AttributeDefinitions: []types.AttributeDefinition{
		{
			AttributeName: aws.String("Address"),
			AttributeType: "S",
		},
		{
			AttributeName: aws.String("Event"),
			AttributeType: "N",
		},
		{
			AttributeName: aws.String("Created"),
			AttributeType: "N",
		},
	},
	KeySchema: []types.KeySchemaElement{
		{
			AttributeName: aws.String("Address"),
			KeyType:       "HASH",
		},
		{
			AttributeName: aws.String("Created"),
			KeyType:       "RANGE",
		},
	},
	GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
		{
			IndexName: aws.String("address"),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("Address"),
					KeyType:       "HASH",
				},
			},
			Projection: &types.Projection{
				// maybe just address...?
				ProjectionType: "ALL",
			},
		},
		{
			IndexName: aws.String("event"),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("Event"),
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
