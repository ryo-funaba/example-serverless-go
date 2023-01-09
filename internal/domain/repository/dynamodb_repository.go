package repository

import dynamoDB "github.com/aws/aws-sdk-go/service/dynamodb"

type DynamodbRepository interface {
	GetItem(input dynamoDB.GetItemInput) (*dynamoDB.GetItemOutput, error)
	PutItem(input dynamoDB.PutItemInput) (*dynamoDB.PutItemOutput, error)
}
