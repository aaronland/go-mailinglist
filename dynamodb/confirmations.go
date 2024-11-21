package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var DynamoDBConfirmationsTable = &dynamodb.CreateTableInput{
	AttributeDefinitions: []types.AttributeDefinition{
		{
			AttributeName: aws.String("Code"),
			AttributeType: "S",
		},
		{
			AttributeName: aws.String("Address"),
			AttributeType: "S",
		},
		{
			AttributeName: aws.String("Created"),
			AttributeType: "N",
		},
	},
	KeySchema: []types.KeySchemaElement{
		{
			AttributeName: aws.String("Code"),
			KeyType:       "HASH",
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
				ProjectionType: "ALL",
			},
		},
		{
			IndexName: aws.String("created"),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("Created"),
					KeyType:       "HASH",
				},
			},
			Projection: &types.Projection{
				ProjectionType: "INCLUDE",
				NonKeyAttributes: []string{
					"code",
				},
			},
		},
	},
	BillingMode: types.BillingModePayPerRequest,
}
