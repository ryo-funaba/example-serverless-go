package dynamodb

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	dynamoDB "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type dynamodbSession struct {
	sess *session.Session
}

type DynamodbSession interface {
	GetClient() dynamodbiface.DynamoDBAPI
}

func NewDynamodbSession() DynamodbSession {
	var (
		dynamodbSess *dynamodbSession
		dynamodbOnce sync.Once
	)

	dynamodbOnce.Do(func() {
		opt := session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}
		sess := session.Must(session.NewSessionWithOptions(opt))

		dynamodbSess = &dynamodbSession{sess}
	})

	return dynamodbSess
}

func (c *dynamodbSession) GetClient() dynamodbiface.DynamoDBAPI {
	svc := dynamoDB.New(c.sess)
	api := dynamodbiface.DynamoDBAPI(svc)

	return api
}
