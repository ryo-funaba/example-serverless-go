package dynamodb

import (
	"fmt"

	dynamoDB "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/ryo-funaba/example_echo/internal/domain/repository"
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
	"github.com/ryo-funaba/example_echo/internal/utils/logutil"
)

const dynamodbRepositoryStructName = "dynamodbRepository"

type dynamodbRepository struct {
	svc dynamodbiface.DynamoDBAPI
}

func NewDynamodbRepository(svc dynamodbiface.DynamoDBAPI) repository.DynamodbRepository {
	return &dynamodbRepository{svc: svc}
}

func (r *dynamodbRepository) GetItem(input dynamoDB.GetItemInput) (*dynamoDB.GetItemOutput, error) {
	logutil.PrintFuncName(dynamodbRepositoryStructName, "GetItem")

	fmt.Printf("GetItem input %+v\n", input)

	obj, err := r.svc.GetItem(&input)
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	fmt.Printf("GetItem output %+v\n", *obj)

	return obj, nil
}

func (r *dynamodbRepository) PutItem(input dynamoDB.PutItemInput) (*dynamoDB.PutItemOutput, error) {
	logutil.PrintFuncName(dynamodbRepositoryStructName, "PutItem")

	fmt.Printf("PutItem input %+v\n", input)

	obj, err := r.svc.PutItem(&input)
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	fmt.Printf("PutItem output %+v\n", *obj)

	return obj, nil
}
