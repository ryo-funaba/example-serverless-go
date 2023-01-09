package repository

import "github.com/aws/aws-sdk-go/service/s3"

type S3Repository interface {
	GetObject(input s3.GetObjectInput) (*s3.GetObjectOutput, error)
	PutObject(input s3.PutObjectInput) (*s3.PutObjectOutput, error)
	DeleteObject(input s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
	CopyObject(input s3.CopyObjectInput) (*s3.CopyObjectOutput, error)
}
