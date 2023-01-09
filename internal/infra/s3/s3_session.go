package s3

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type s3Session struct {
	Session *session.Session
}

type S3Session interface {
	GetClient() s3iface.S3API
}

func NewS3Session(accessKey, secretAccessKey, region string) S3Session {
	var (
		s3Sess *s3Session
		s3Once sync.Once
	)

	s3Once.Do(func() {
		creds := credentials.NewStaticCredentials(accessKey, secretAccessKey, "")
		conf := &aws.Config{
			Credentials: creds,
			Region:      aws.String(region),
		}
		sess := session.Must(session.NewSession(conf))
		s3Sess = &s3Session{sess}
	})

	return s3Sess
}

func (c *s3Session) GetClient() s3iface.S3API {
	svc := s3.New(c.Session)
	api := s3iface.S3API(svc)

	return api
}
