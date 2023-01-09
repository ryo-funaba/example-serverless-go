package s3

import (
	"fmt"

	"github.com/ryo-funaba/example_echo/internal/domain/repository"
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
	"github.com/ryo-funaba/example_echo/internal/utils/logutil"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

const s3RepositoryStructName = "s3Repository"

type s3Repository struct {
	svc s3iface.S3API
}

func NewS3Repository(svc s3iface.S3API) repository.S3Repository {
	return &s3Repository{svc: svc}
}

func (r *s3Repository) GetObject(input s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	logutil.PrintFuncName(s3RepositoryStructName, "GetObject")

	fmt.Printf("GetObject input %+v\n", input)

	obj, err := r.svc.GetObject(&input)
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	fmt.Printf("GetObject output %+v\n", *obj)

	return obj, nil
}

func (r *s3Repository) PutObject(input s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	logutil.PrintFuncName(s3RepositoryStructName, "PutObject")

	fmt.Printf("PutObject input %+v\n", input)

	obj, err := r.svc.PutObject(&input)
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	fmt.Printf("PutObject output %+v\n", *obj)

	return obj, nil
}

func (r *s3Repository) DeleteObject(input s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	logutil.PrintFuncName(s3RepositoryStructName, "DeleteObject")

	fmt.Printf("DeleteObject input %+v\n", input)

	obj, err := r.svc.DeleteObject(&input)
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	fmt.Printf("DeleteObject output %+v\n", *obj)

	return obj, nil
}

func (r *s3Repository) CopyObject(input s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	logutil.PrintFuncName(s3RepositoryStructName, "CopyObject")

	fmt.Printf("CopyObject input %+v\n", input)

	obj, err := r.svc.CopyObject(&input)
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	fmt.Printf("CopyObject output %+v\n", *obj)

	return obj, nil
}
