package dynamodb

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	dynamoDB "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/go-cmp/cmp"
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
	"google.golang.org/protobuf/testing/protocmp"
)

func Test_dynamodbRepository_GetItem(t *testing.T) {
	const (
		table        = "sample_table"
		pk           = "sample_pk"
		sk           = "sample_sk"
		invalidTable = "invalid_table"
		invalidPK    = "invalid_pk"
		invalidSK    = "invalid_sk"
	)

	type args struct {
		input dynamoDB.GetItemInput
	}

	tests := []struct {
		name    string
		svc     func(t *testing.T) dynamodbiface.DynamoDBAPI
		args    func() args
		want    func() *dynamoDB.GetItemOutput
		wantErr error
	}{
		{
			name: "存在するテーブルとPKとSKを指定した場合、データが1件返される",
			svc: func(t *testing.T) dynamodbiface.DynamoDBAPI {
				return mockDynamodbRepository{Table: table, PK: pk, SK: sk}
			},
			args: func() args {
				return args{
					input: dynamoDB.GetItemInput{
						Key: map[string]*dynamoDB.AttributeValue{
							"PK": {
								S: aws.String(pk),
							},
							"SK": {
								S: aws.String(sk),
							},
						},
						TableName: aws.String(table),
					},
				}
			},
			want: func() *dynamoDB.GetItemOutput {
				return &dynamoDB.GetItemOutput{}
			},
			wantErr: nil,
		},
		{
			name: "存在しないテーブルを指定した場合、エラーが返される",
			svc: func(t *testing.T) dynamodbiface.DynamoDBAPI {
				return mockDynamodbRepository{Table: table, PK: pk, SK: sk}
			},
			args: func() args {
				return args{
					input: dynamoDB.GetItemInput{
						Key: map[string]*dynamoDB.AttributeValue{
							"PK": {
								S: aws.String(pk),
							},
							"SK": {
								S: aws.String(sk),
							},
						},
						TableName: aws.String(invalidTable),
					},
				}
			},
			want: func() *dynamoDB.GetItemOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "TableName is invalid"),
		},
		{
			name: "存在しないPKを指定した場合、エラーが返される",
			svc: func(t *testing.T) dynamodbiface.DynamoDBAPI {
				return mockDynamodbRepository{Table: table, PK: pk, SK: sk}
			},
			args: func() args {
				return args{
					input: dynamoDB.GetItemInput{
						Key: map[string]*dynamoDB.AttributeValue{
							"PK": {
								S: aws.String(invalidPK),
							},
							"SK": {
								S: aws.String(sk),
							},
						},
						TableName: aws.String(table),
					},
				}
			},
			want: func() *dynamoDB.GetItemOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "PK is invalid"),
		},
		{
			name: "存在しないSKを指定した場合、エラーが返される",
			svc: func(t *testing.T) dynamodbiface.DynamoDBAPI {
				return mockDynamodbRepository{Table: table, PK: pk, SK: sk}
			},
			args: func() args {
				return args{
					input: dynamoDB.GetItemInput{
						Key: map[string]*dynamoDB.AttributeValue{
							"PK": {
								S: aws.String(pk),
							},
							"SK": {
								S: aws.String(invalidSK),
							},
						},
						TableName: aws.String(table),
					},
				}
			},
			want: func() *dynamoDB.GetItemOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "SK is invalid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDynamodbRepository(tt.svc(t))

			got, err := r.GetItem(tt.args().input)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("dynamodbRepository.GetItem() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(got, tt.want(), protocmp.Transform()); diff != "" {
				t.Errorf("Failed due to %s", diff)
			}
		})
	}
}

func Test_dynamodbRepository_PutItem(t *testing.T) {
	const (
		table        = "sample_table"
		pk           = "sample_pk"
		sk           = "sample_sk"
		invalidTable = "invalid_table"
		invalidPK    = "invalid_pk"
		invalidSK    = "invalid_sk"
	)

	type args struct {
		input dynamoDB.PutItemInput
	}

	tests := []struct {
		name    string
		svc     func(t *testing.T) dynamodbiface.DynamoDBAPI
		args    func() args
		want    func() *dynamoDB.PutItemOutput
		wantErr error
	}{
		{
			name: "存在するテーブルとPKとSKを指定した場合、データが1件返される",
			svc: func(t *testing.T) dynamodbiface.DynamoDBAPI {
				return mockDynamodbRepository{Table: table, PK: pk, SK: sk}
			},
			args: func() args {
				item := map[string]interface{}{
					"PK": pk,
					"SK": sk,
				}

				av, err := dynamodbattribute.MarshalMap(item)
				if err != nil {
					panic(err)
				}

				input := dynamodb.PutItemInput{
					Item:      av,
					TableName: aws.String(table),
				}

				return args{input: input}
			},
			want: func() *dynamoDB.PutItemOutput {
				return &dynamoDB.PutItemOutput{}
			},
			wantErr: nil,
		},
		{
			name: "存在しないテーブルを指定した場合、エラーが返される",
			svc: func(t *testing.T) dynamodbiface.DynamoDBAPI {
				return mockDynamodbRepository{Table: table, PK: pk, SK: sk}
			},
			args: func() args {
				item := map[string]interface{}{
					"PK": pk,
					"SK": sk,
				}

				av, err := dynamodbattribute.MarshalMap(item)
				if err != nil {
					panic(err)
				}

				input := dynamodb.PutItemInput{
					Item:      av,
					TableName: aws.String(invalidTable),
				}

				return args{input: input}
			},
			want: func() *dynamoDB.PutItemOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "TableName is invalid"),
		},
		{
			name: "存在しないPKを指定した場合、エラーが返される",
			svc: func(t *testing.T) dynamodbiface.DynamoDBAPI {
				return mockDynamodbRepository{Table: table, PK: pk, SK: sk}
			},
			args: func() args {
				item := map[string]interface{}{
					"PK": invalidPK,
					"SK": sk,
				}

				av, err := dynamodbattribute.MarshalMap(item)
				if err != nil {
					panic(err)
				}

				input := dynamodb.PutItemInput{
					Item:      av,
					TableName: aws.String(table),
				}

				return args{input: input}
			},
			want: func() *dynamoDB.PutItemOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "PK is invalid"),
		},
		{
			name: "存在しないSKを指定した場合、エラーが返される",
			svc: func(t *testing.T) dynamodbiface.DynamoDBAPI {
				return mockDynamodbRepository{Table: table, PK: pk, SK: sk}
			},
			args: func() args {
				item := map[string]interface{}{
					"PK": pk,
					"SK": invalidSK,
				}

				av, err := dynamodbattribute.MarshalMap(item)
				if err != nil {
					panic(err)
				}

				input := dynamodb.PutItemInput{
					Item:      av,
					TableName: aws.String(table),
				}

				return args{input: input}
			},
			want: func() *dynamoDB.PutItemOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "SK is invalid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDynamodbRepository(tt.svc(t))

			got, err := r.PutItem(tt.args().input)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("dynamodbRepository.PutItem() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(got, tt.want(), protocmp.Transform()); diff != "" {
				t.Errorf("Failed due to %s", diff)
			}
		})
	}
}
