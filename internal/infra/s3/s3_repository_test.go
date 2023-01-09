package s3

import (
	"testing"

	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/google/go-cmp/cmp"
)

func Test_s3Repository_GetObject(t *testing.T) {
	const (
		bucket        = "sample_bucket"
		key           = "sample_key"
		invalidBucket = "invalid_bucket"
		invalidKey    = "invalid_key"
	)

	type args struct {
		input s3.GetObjectInput
	}

	tests := []struct {
		name    string
		svc     func(t *testing.T) s3iface.S3API
		args    func() args
		want    func() *s3.GetObjectOutput
		wantErr error
	}{
		{
			name: "存在するBucketとKeyを指定した場合、データが1件返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.GetObjectInput{
						Bucket: aws.String(bucket),
						Key:    aws.String(key),
					},
				}
			},
			want: func() *s3.GetObjectOutput {
				return &s3.GetObjectOutput{}
			},
			wantErr: nil,
		},
		{
			name: "存在しないBucketを指定した場合、エラーが返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.GetObjectInput{
						Bucket: aws.String(invalidBucket),
						Key:    aws.String(key),
					},
				}
			},
			want: func() *s3.GetObjectOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "bucket name is not valid"),
		},
		{
			name: "存在しないKeyを指定した場合、エラーが返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.GetObjectInput{
						Bucket: aws.String(bucket),
						Key:    aws.String(invalidKey),
					},
				}
			},
			want: func() *s3.GetObjectOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "key name is not valid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewS3Repository(tt.svc(t))

			got, err := r.GetObject(tt.args().input)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("s3Repository.GetObject() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(got, tt.want()); diff != "" {
				t.Errorf("Failed due to %s", diff)
			}
		})
	}
}

func Test_s3Repository_PutObject(t *testing.T) {
	const (
		bucket        = "sample_bucket"
		key           = "sample_key"
		invalidBucket = "invalid_bucket"
		invalidKey    = "invalid_key"
	)

	type args struct {
		input s3.PutObjectInput
	}

	tests := []struct {
		name    string
		svc     func(t *testing.T) s3iface.S3API
		args    func() args
		want    func() *s3.PutObjectOutput
		wantErr error
	}{
		{
			name: "存在するBucketとKeyを指定した場合、データが1件返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.PutObjectInput{
						Bucket: aws.String(bucket),
						Key:    aws.String(key),
					},
				}
			},
			want: func() *s3.PutObjectOutput {
				return &s3.PutObjectOutput{}
			},
			wantErr: nil,
		},
		{
			name: "存在しないBucketを指定した場合、エラーが返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.PutObjectInput{
						Bucket: aws.String(invalidBucket),
						Key:    aws.String(key),
					},
				}
			},
			want: func() *s3.PutObjectOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "bucket name is not valid"),
		},
		{
			name: "存在しないKeyを指定した場合、エラーが返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.PutObjectInput{
						Bucket: aws.String(bucket),
						Key:    aws.String(invalidKey),
					},
				}
			},
			want: func() *s3.PutObjectOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "key name is not valid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewS3Repository(tt.svc(t))

			got, err := r.PutObject(tt.args().input)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("s3Repository.PutObject() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(got, tt.want()); diff != "" {
				t.Errorf("Failed due to %s", diff)
			}
		})
	}
}

func Test_s3Repository_DeleteObject(t *testing.T) {
	const (
		bucket        = "sample_bucket"
		key           = "sample_key"
		invalidBucket = "invalid_bucket"
		invalidKey    = "invalid_key"
	)

	type args struct {
		input s3.DeleteObjectInput
	}

	tests := []struct {
		name    string
		svc     func(t *testing.T) s3iface.S3API
		args    func() args
		want    func() *s3.DeleteObjectOutput
		wantErr error
	}{
		{
			name: "存在するBucketとKeyを指定した場合、データが1件返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.DeleteObjectInput{
						Bucket: aws.String(bucket),
						Key:    aws.String(key),
					},
				}
			},
			want: func() *s3.DeleteObjectOutput {
				return &s3.DeleteObjectOutput{}
			},
			wantErr: nil,
		},
		{
			name: "存在しないBucketを指定した場合、エラーが返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.DeleteObjectInput{
						Bucket: aws.String(invalidBucket),
						Key:    aws.String(key),
					},
				}
			},
			want: func() *s3.DeleteObjectOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "bucket name is not valid"),
		},
		{
			name: "存在しないKeyを指定した場合、エラーが返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.DeleteObjectInput{
						Bucket: aws.String(bucket),
						Key:    aws.String(invalidKey),
					},
				}
			},
			want: func() *s3.DeleteObjectOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "key name is not valid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewS3Repository(tt.svc(t))

			got, err := r.DeleteObject(tt.args().input)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("s3Repository.DeleteObject() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(got, tt.want()); diff != "" {
				t.Errorf("Failed due to %s", diff)
			}
		})
	}
}

func Test_s3Repository_CopyObject(t *testing.T) {
	const (
		bucket        = "sample_bucket"
		key           = "sample_key"
		invalidBucket = "invalid_bucket"
		invalidKey    = "invalid_key"
	)

	type args struct {
		input s3.CopyObjectInput
	}

	tests := []struct {
		name    string
		svc     func(t *testing.T) s3iface.S3API
		args    func() args
		want    func() *s3.CopyObjectOutput
		wantErr error
	}{
		{
			name: "存在するBucketとKeyを指定した場合、データが1件返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.CopyObjectInput{
						Bucket: aws.String(bucket),
						Key:    aws.String(key),
					},
				}
			},
			want: func() *s3.CopyObjectOutput {
				return &s3.CopyObjectOutput{}
			},
			wantErr: nil,
		},
		{
			name: "存在しないBucketを指定した場合、エラーが返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.CopyObjectInput{
						Bucket: aws.String(invalidBucket),
						Key:    aws.String(key),
					},
				}
			},
			want: func() *s3.CopyObjectOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "bucket name is not valid"),
		},
		{
			name: "存在しないKeyを指定した場合、エラーが返される",
			svc: func(t *testing.T) s3iface.S3API {
				return mockS3Repository{bucket: bucket, key: key}
			},
			args: func() args {
				return args{
					input: s3.CopyObjectInput{
						Bucket: aws.String(bucket),
						Key:    aws.String(invalidKey),
					},
				}
			},
			want: func() *s3.CopyObjectOutput {
				return nil
			},
			wantErr: errorutil.Errorf(errorutil.Unknown, "key name is not valid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewS3Repository(tt.svc(t))

			got, err := r.CopyObject(tt.args().input)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("s3Repository.CopyObject() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(got, tt.want()); diff != "" {
				t.Errorf("Failed due to %s", diff)
			}
		})
	}
}
