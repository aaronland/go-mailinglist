package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var DynamoDBEventLogsTable = &dynamodb.CreateTableInput{
	AttributeDefinitions: []types.AttributeDefinition{
		{
			AttributeName: aws.String("address"),
			AttributeType: "S",
		},
		{
			AttributeName: aws.String("event"),
			AttributeType: "N",
		},
		{
			AttributeName: aws.String("created"),
			AttributeType: "N",
		},
	},
	KeySchema: []types.KeySchemaElement{
		{
			AttributeName: aws.String("address"),
			KeyType:       "HASH",
		},
		{
			AttributeName: aws.String("created"),
			KeyType:       "RANGE",
		},
	},
	GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
		{
			IndexName: aws.String("address"),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("address"),
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
					AttributeName: aws.String("event"),
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
