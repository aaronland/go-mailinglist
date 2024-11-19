package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var DynamoDBConfirmationsTable = &dynamodb.CreateTableInput{
	AttributeDefinitions: []types.AttributeDefinition{
		{
			AttributeName: aws.String("code"),
			AttributeType: "S",
		},
		{
			AttributeName: aws.String("address"),
			AttributeType: "S",
		},
		{
			AttributeName: aws.String("created"),
			AttributeType: "N",
		},
	},
	KeySchema: []types.KeySchemaElement{
		{
			AttributeName: aws.String("code"),
			KeyType:       "HASH",
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
				ProjectionType: "ALL",
			},
		},
		{
			IndexName: aws.String("created"),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("created"),
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
