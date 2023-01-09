package s3

import (
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
)

type mockS3Repository struct {
	bucket string
	key    string
}

func (m mockS3Repository) AbortMultipartUpload(*s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	return nil, nil
}

func (m mockS3Repository) AbortMultipartUploadWithContext(aws.Context, *s3.AbortMultipartUploadInput, ...request.Option) (*s3.AbortMultipartUploadOutput, error) {
	return nil, nil
}

func (m mockS3Repository) AbortMultipartUploadRequest(*s3.AbortMultipartUploadInput) (*request.Request, *s3.AbortMultipartUploadOutput) {
	return nil, nil
}

func (m mockS3Repository) CompleteMultipartUpload(*s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	return nil, nil
}

func (m mockS3Repository) CompleteMultipartUploadWithContext(aws.Context, *s3.CompleteMultipartUploadInput, ...request.Option) (*s3.CompleteMultipartUploadOutput, error) {
	return nil, nil
}

func (m mockS3Repository) CompleteMultipartUploadRequest(*s3.CompleteMultipartUploadInput) (*request.Request, *s3.CompleteMultipartUploadOutput) {
	return nil, nil
}

func (m mockS3Repository) CopyObject(input *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	if *input.Bucket != m.bucket {
		return nil, errorutil.Errorf(errorutil.Unknown, "bucket name is not valid")
	}

	if *input.Key != m.key {
		return nil, errorutil.Errorf(errorutil.Unknown, "key name is not valid")
	}

	return &s3.CopyObjectOutput{}, nil
}

func (m mockS3Repository) CopyObjectWithContext(aws.Context, *s3.CopyObjectInput, ...request.Option) (*s3.CopyObjectOutput, error) {
	return nil, nil
}

func (m mockS3Repository) CopyObjectRequest(*s3.CopyObjectInput) (*request.Request, *s3.CopyObjectOutput) {
	return nil, nil
}

func (m mockS3Repository) CreateBucket(*s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	return nil, nil
}

func (m mockS3Repository) CreateBucketWithContext(aws.Context, *s3.CreateBucketInput, ...request.Option) (*s3.CreateBucketOutput, error) {
	return nil, nil
}

func (m mockS3Repository) CreateBucketRequest(*s3.CreateBucketInput) (*request.Request, *s3.CreateBucketOutput) {
	return nil, nil
}

func (m mockS3Repository) CreateMultipartUpload(*s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	return nil, nil
}

func (m mockS3Repository) CreateMultipartUploadWithContext(aws.Context, *s3.CreateMultipartUploadInput, ...request.Option) (*s3.CreateMultipartUploadOutput, error) {
	return nil, nil
}

func (m mockS3Repository) CreateMultipartUploadRequest(*s3.CreateMultipartUploadInput) (*request.Request, *s3.CreateMultipartUploadOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucket(*s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketWithContext(aws.Context, *s3.DeleteBucketInput, ...request.Option) (*s3.DeleteBucketOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketRequest(*s3.DeleteBucketInput) (*request.Request, *s3.DeleteBucketOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketAnalyticsConfiguration(*s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketAnalyticsConfigurationWithContext(aws.Context, *s3.DeleteBucketAnalyticsConfigurationInput, ...request.Option) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketAnalyticsConfigurationRequest(*s3.DeleteBucketAnalyticsConfigurationInput) (*request.Request, *s3.DeleteBucketAnalyticsConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketCors(*s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketCorsWithContext(aws.Context, *s3.DeleteBucketCorsInput, ...request.Option) (*s3.DeleteBucketCorsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketCorsRequest(*s3.DeleteBucketCorsInput) (*request.Request, *s3.DeleteBucketCorsOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketEncryption(*s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketEncryptionWithContext(aws.Context, *s3.DeleteBucketEncryptionInput, ...request.Option) (*s3.DeleteBucketEncryptionOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketEncryptionRequest(*s3.DeleteBucketEncryptionInput) (*request.Request, *s3.DeleteBucketEncryptionOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketIntelligentTieringConfiguration(*s3.DeleteBucketIntelligentTieringConfigurationInput) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketIntelligentTieringConfigurationWithContext(aws.Context, *s3.DeleteBucketIntelligentTieringConfigurationInput, ...request.Option) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketIntelligentTieringConfigurationRequest(*s3.DeleteBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.DeleteBucketIntelligentTieringConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketInventoryConfiguration(*s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketInventoryConfigurationWithContext(aws.Context, *s3.DeleteBucketInventoryConfigurationInput, ...request.Option) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketInventoryConfigurationRequest(*s3.DeleteBucketInventoryConfigurationInput) (*request.Request, *s3.DeleteBucketInventoryConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketLifecycle(*s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketLifecycleWithContext(aws.Context, *s3.DeleteBucketLifecycleInput, ...request.Option) (*s3.DeleteBucketLifecycleOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketLifecycleRequest(*s3.DeleteBucketLifecycleInput) (*request.Request, *s3.DeleteBucketLifecycleOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketMetricsConfiguration(*s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketMetricsConfigurationWithContext(aws.Context, *s3.DeleteBucketMetricsConfigurationInput, ...request.Option) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketMetricsConfigurationRequest(*s3.DeleteBucketMetricsConfigurationInput) (*request.Request, *s3.DeleteBucketMetricsConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketOwnershipControls(*s3.DeleteBucketOwnershipControlsInput) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketOwnershipControlsWithContext(aws.Context, *s3.DeleteBucketOwnershipControlsInput, ...request.Option) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketOwnershipControlsRequest(*s3.DeleteBucketOwnershipControlsInput) (*request.Request, *s3.DeleteBucketOwnershipControlsOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketPolicy(*s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketPolicyWithContext(aws.Context, *s3.DeleteBucketPolicyInput, ...request.Option) (*s3.DeleteBucketPolicyOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketPolicyRequest(*s3.DeleteBucketPolicyInput) (*request.Request, *s3.DeleteBucketPolicyOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketReplication(*s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketReplicationWithContext(aws.Context, *s3.DeleteBucketReplicationInput, ...request.Option) (*s3.DeleteBucketReplicationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketReplicationRequest(*s3.DeleteBucketReplicationInput) (*request.Request, *s3.DeleteBucketReplicationOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketTagging(*s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketTaggingWithContext(aws.Context, *s3.DeleteBucketTaggingInput, ...request.Option) (*s3.DeleteBucketTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketTaggingRequest(*s3.DeleteBucketTaggingInput) (*request.Request, *s3.DeleteBucketTaggingOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketWebsite(*s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketWebsiteWithContext(aws.Context, *s3.DeleteBucketWebsiteInput, ...request.Option) (*s3.DeleteBucketWebsiteOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteBucketWebsiteRequest(*s3.DeleteBucketWebsiteInput) (*request.Request, *s3.DeleteBucketWebsiteOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	if *input.Bucket != m.bucket {
		return nil, errorutil.Errorf(errorutil.Unknown, "bucket name is not valid")
	}

	if *input.Key != m.key {
		return nil, errorutil.Errorf(errorutil.Unknown, "key name is not valid")
	}

	return &s3.DeleteObjectOutput{}, nil
}

func (m mockS3Repository) DeleteObjectWithContext(aws.Context, *s3.DeleteObjectInput, ...request.Option) (*s3.DeleteObjectOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteObjectRequest(*s3.DeleteObjectInput) (*request.Request, *s3.DeleteObjectOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteObjectTagging(*s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteObjectTaggingWithContext(aws.Context, *s3.DeleteObjectTaggingInput, ...request.Option) (*s3.DeleteObjectTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteObjectTaggingRequest(*s3.DeleteObjectTaggingInput) (*request.Request, *s3.DeleteObjectTaggingOutput) {
	return nil, nil
}

func (m mockS3Repository) DeleteObjects(*s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteObjectsWithContext(aws.Context, *s3.DeleteObjectsInput, ...request.Option) (*s3.DeleteObjectsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeleteObjectsRequest(*s3.DeleteObjectsInput) (*request.Request, *s3.DeleteObjectsOutput) {
	return nil, nil
}

func (m mockS3Repository) DeletePublicAccessBlock(*s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeletePublicAccessBlockWithContext(aws.Context, *s3.DeletePublicAccessBlockInput, ...request.Option) (*s3.DeletePublicAccessBlockOutput, error) {
	return nil, nil
}

func (m mockS3Repository) DeletePublicAccessBlockRequest(*s3.DeletePublicAccessBlockInput) (*request.Request, *s3.DeletePublicAccessBlockOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketAccelerateConfiguration(*s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketAccelerateConfigurationWithContext(aws.Context, *s3.GetBucketAccelerateConfigurationInput, ...request.Option) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketAccelerateConfigurationRequest(*s3.GetBucketAccelerateConfigurationInput) (*request.Request, *s3.GetBucketAccelerateConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketAcl(*s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketAclWithContext(aws.Context, *s3.GetBucketAclInput, ...request.Option) (*s3.GetBucketAclOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketAclRequest(*s3.GetBucketAclInput) (*request.Request, *s3.GetBucketAclOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketAnalyticsConfiguration(*s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketAnalyticsConfigurationWithContext(aws.Context, *s3.GetBucketAnalyticsConfigurationInput, ...request.Option) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketAnalyticsConfigurationRequest(*s3.GetBucketAnalyticsConfigurationInput) (*request.Request, *s3.GetBucketAnalyticsConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketCors(*s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketCorsWithContext(aws.Context, *s3.GetBucketCorsInput, ...request.Option) (*s3.GetBucketCorsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketCorsRequest(*s3.GetBucketCorsInput) (*request.Request, *s3.GetBucketCorsOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketEncryption(*s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketEncryptionWithContext(aws.Context, *s3.GetBucketEncryptionInput, ...request.Option) (*s3.GetBucketEncryptionOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketEncryptionRequest(*s3.GetBucketEncryptionInput) (*request.Request, *s3.GetBucketEncryptionOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketIntelligentTieringConfiguration(*s3.GetBucketIntelligentTieringConfigurationInput) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketIntelligentTieringConfigurationWithContext(aws.Context, *s3.GetBucketIntelligentTieringConfigurationInput, ...request.Option) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketIntelligentTieringConfigurationRequest(*s3.GetBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.GetBucketIntelligentTieringConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketInventoryConfiguration(*s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketInventoryConfigurationWithContext(aws.Context, *s3.GetBucketInventoryConfigurationInput, ...request.Option) (*s3.GetBucketInventoryConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketInventoryConfigurationRequest(*s3.GetBucketInventoryConfigurationInput) (*request.Request, *s3.GetBucketInventoryConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLifecycle(*s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLifecycleWithContext(aws.Context, *s3.GetBucketLifecycleInput, ...request.Option) (*s3.GetBucketLifecycleOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLifecycleRequest(*s3.GetBucketLifecycleInput) (*request.Request, *s3.GetBucketLifecycleOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLifecycleConfiguration(*s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLifecycleConfigurationWithContext(aws.Context, *s3.GetBucketLifecycleConfigurationInput, ...request.Option) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLifecycleConfigurationRequest(*s3.GetBucketLifecycleConfigurationInput) (*request.Request, *s3.GetBucketLifecycleConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLocation(*s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLocationWithContext(aws.Context, *s3.GetBucketLocationInput, ...request.Option) (*s3.GetBucketLocationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLocationRequest(*s3.GetBucketLocationInput) (*request.Request, *s3.GetBucketLocationOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLogging(*s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLoggingWithContext(aws.Context, *s3.GetBucketLoggingInput, ...request.Option) (*s3.GetBucketLoggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketLoggingRequest(*s3.GetBucketLoggingInput) (*request.Request, *s3.GetBucketLoggingOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketMetricsConfiguration(*s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketMetricsConfigurationWithContext(aws.Context, *s3.GetBucketMetricsConfigurationInput, ...request.Option) (*s3.GetBucketMetricsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketMetricsConfigurationRequest(*s3.GetBucketMetricsConfigurationInput) (*request.Request, *s3.GetBucketMetricsConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketNotification(*s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketNotificationWithContext(aws.Context, *s3.GetBucketNotificationConfigurationRequest, ...request.Option) (*s3.NotificationConfigurationDeprecated, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketNotificationRequest(*s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfigurationDeprecated) {
	return nil, nil
}

func (m mockS3Repository) GetBucketNotificationConfiguration(*s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketNotificationConfigurationWithContext(aws.Context, *s3.GetBucketNotificationConfigurationRequest, ...request.Option) (*s3.NotificationConfiguration, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketNotificationConfigurationRequest(*s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfiguration) {
	return nil, nil
}

func (m mockS3Repository) GetBucketOwnershipControls(*s3.GetBucketOwnershipControlsInput) (*s3.GetBucketOwnershipControlsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketOwnershipControlsWithContext(aws.Context, *s3.GetBucketOwnershipControlsInput, ...request.Option) (*s3.GetBucketOwnershipControlsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketOwnershipControlsRequest(*s3.GetBucketOwnershipControlsInput) (*request.Request, *s3.GetBucketOwnershipControlsOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketPolicy(*s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketPolicyWithContext(aws.Context, *s3.GetBucketPolicyInput, ...request.Option) (*s3.GetBucketPolicyOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketPolicyRequest(*s3.GetBucketPolicyInput) (*request.Request, *s3.GetBucketPolicyOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketPolicyStatus(*s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketPolicyStatusWithContext(aws.Context, *s3.GetBucketPolicyStatusInput, ...request.Option) (*s3.GetBucketPolicyStatusOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketPolicyStatusRequest(*s3.GetBucketPolicyStatusInput) (*request.Request, *s3.GetBucketPolicyStatusOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketReplication(*s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketReplicationWithContext(aws.Context, *s3.GetBucketReplicationInput, ...request.Option) (*s3.GetBucketReplicationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketReplicationRequest(*s3.GetBucketReplicationInput) (*request.Request, *s3.GetBucketReplicationOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketRequestPayment(*s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketRequestPaymentWithContext(aws.Context, *s3.GetBucketRequestPaymentInput, ...request.Option) (*s3.GetBucketRequestPaymentOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketRequestPaymentRequest(*s3.GetBucketRequestPaymentInput) (*request.Request, *s3.GetBucketRequestPaymentOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketTagging(*s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketTaggingWithContext(aws.Context, *s3.GetBucketTaggingInput, ...request.Option) (*s3.GetBucketTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketTaggingRequest(*s3.GetBucketTaggingInput) (*request.Request, *s3.GetBucketTaggingOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketVersioning(*s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketVersioningWithContext(aws.Context, *s3.GetBucketVersioningInput, ...request.Option) (*s3.GetBucketVersioningOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketVersioningRequest(*s3.GetBucketVersioningInput) (*request.Request, *s3.GetBucketVersioningOutput) {
	return nil, nil
}

func (m mockS3Repository) GetBucketWebsite(*s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketWebsiteWithContext(aws.Context, *s3.GetBucketWebsiteInput, ...request.Option) (*s3.GetBucketWebsiteOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetBucketWebsiteRequest(*s3.GetBucketWebsiteInput) (*request.Request, *s3.GetBucketWebsiteOutput) {
	return nil, nil
}

func (m mockS3Repository) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	if *input.Bucket != m.bucket {
		return nil, errorutil.Errorf(errorutil.Unknown, "bucket name is not valid")
	}

	if *input.Key != m.key {
		return nil, errorutil.Errorf(errorutil.Unknown, "key name is not valid")
	}

	return &s3.GetObjectOutput{}, nil
}

func (m mockS3Repository) GetObjectWithContext(aws.Context, *s3.GetObjectInput, ...request.Option) (*s3.GetObjectOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectRequest(*s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput) {
	return nil, nil
}

func (m mockS3Repository) GetObjectAcl(*s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectAclWithContext(aws.Context, *s3.GetObjectAclInput, ...request.Option) (*s3.GetObjectAclOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectAclRequest(*s3.GetObjectAclInput) (*request.Request, *s3.GetObjectAclOutput) {
	return nil, nil
}

func (m mockS3Repository) GetObjectAttributes(*s3.GetObjectAttributesInput) (*s3.GetObjectAttributesOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectAttributesWithContext(aws.Context, *s3.GetObjectAttributesInput, ...request.Option) (*s3.GetObjectAttributesOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectAttributesRequest(*s3.GetObjectAttributesInput) (*request.Request, *s3.GetObjectAttributesOutput) {
	return nil, nil
}

func (m mockS3Repository) GetObjectLegalHold(*s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectLegalHoldWithContext(aws.Context, *s3.GetObjectLegalHoldInput, ...request.Option) (*s3.GetObjectLegalHoldOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectLegalHoldRequest(*s3.GetObjectLegalHoldInput) (*request.Request, *s3.GetObjectLegalHoldOutput) {
	return nil, nil
}

func (m mockS3Repository) GetObjectLockConfiguration(*s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectLockConfigurationWithContext(aws.Context, *s3.GetObjectLockConfigurationInput, ...request.Option) (*s3.GetObjectLockConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectLockConfigurationRequest(*s3.GetObjectLockConfigurationInput) (*request.Request, *s3.GetObjectLockConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) GetObjectRetention(*s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectRetentionWithContext(aws.Context, *s3.GetObjectRetentionInput, ...request.Option) (*s3.GetObjectRetentionOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectRetentionRequest(*s3.GetObjectRetentionInput) (*request.Request, *s3.GetObjectRetentionOutput) {
	return nil, nil
}

func (m mockS3Repository) GetObjectTagging(*s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectTaggingWithContext(aws.Context, *s3.GetObjectTaggingInput, ...request.Option) (*s3.GetObjectTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectTaggingRequest(*s3.GetObjectTaggingInput) (*request.Request, *s3.GetObjectTaggingOutput) {
	return nil, nil
}

func (m mockS3Repository) GetObjectTorrent(*s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectTorrentWithContext(aws.Context, *s3.GetObjectTorrentInput, ...request.Option) (*s3.GetObjectTorrentOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetObjectTorrentRequest(*s3.GetObjectTorrentInput) (*request.Request, *s3.GetObjectTorrentOutput) {
	return nil, nil
}

func (m mockS3Repository) GetPublicAccessBlock(*s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetPublicAccessBlockWithContext(aws.Context, *s3.GetPublicAccessBlockInput, ...request.Option) (*s3.GetPublicAccessBlockOutput, error) {
	return nil, nil
}

func (m mockS3Repository) GetPublicAccessBlockRequest(*s3.GetPublicAccessBlockInput) (*request.Request, *s3.GetPublicAccessBlockOutput) {
	return nil, nil
}

func (m mockS3Repository) HeadBucket(*s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	return nil, nil
}

func (m mockS3Repository) HeadBucketWithContext(aws.Context, *s3.HeadBucketInput, ...request.Option) (*s3.HeadBucketOutput, error) {
	return nil, nil
}

func (m mockS3Repository) HeadBucketRequest(*s3.HeadBucketInput) (*request.Request, *s3.HeadBucketOutput) {
	return nil, nil
}

func (m mockS3Repository) HeadObject(*s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	return nil, nil
}

func (m mockS3Repository) HeadObjectWithContext(aws.Context, *s3.HeadObjectInput, ...request.Option) (*s3.HeadObjectOutput, error) {
	return nil, nil
}

func (m mockS3Repository) HeadObjectRequest(*s3.HeadObjectInput) (*request.Request, *s3.HeadObjectOutput) {
	return nil, nil
}

func (m mockS3Repository) ListBucketAnalyticsConfigurations(*s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListBucketAnalyticsConfigurationsWithContext(aws.Context, *s3.ListBucketAnalyticsConfigurationsInput, ...request.Option) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListBucketAnalyticsConfigurationsRequest(*s3.ListBucketAnalyticsConfigurationsInput) (*request.Request, *s3.ListBucketAnalyticsConfigurationsOutput) {
	return nil, nil
}

func (m mockS3Repository) ListBucketIntelligentTieringConfigurations(*s3.ListBucketIntelligentTieringConfigurationsInput) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListBucketIntelligentTieringConfigurationsWithContext(aws.Context, *s3.ListBucketIntelligentTieringConfigurationsInput, ...request.Option) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListBucketIntelligentTieringConfigurationsRequest(*s3.ListBucketIntelligentTieringConfigurationsInput) (*request.Request, *s3.ListBucketIntelligentTieringConfigurationsOutput) {
	return nil, nil
}

func (m mockS3Repository) ListBucketInventoryConfigurations(*s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListBucketInventoryConfigurationsWithContext(aws.Context, *s3.ListBucketInventoryConfigurationsInput, ...request.Option) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListBucketInventoryConfigurationsRequest(*s3.ListBucketInventoryConfigurationsInput) (*request.Request, *s3.ListBucketInventoryConfigurationsOutput) {
	return nil, nil
}

func (m mockS3Repository) ListBucketMetricsConfigurations(*s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListBucketMetricsConfigurationsWithContext(aws.Context, *s3.ListBucketMetricsConfigurationsInput, ...request.Option) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListBucketMetricsConfigurationsRequest(*s3.ListBucketMetricsConfigurationsInput) (*request.Request, *s3.ListBucketMetricsConfigurationsOutput) {
	return nil, nil
}

func (m mockS3Repository) ListBuckets(*s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListBucketsWithContext(aws.Context, *s3.ListBucketsInput, ...request.Option) (*s3.ListBucketsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListBucketsRequest(*s3.ListBucketsInput) (*request.Request, *s3.ListBucketsOutput) {
	return nil, nil
}

func (m mockS3Repository) ListMultipartUploads(*s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListMultipartUploadsWithContext(aws.Context, *s3.ListMultipartUploadsInput, ...request.Option) (*s3.ListMultipartUploadsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListMultipartUploadsRequest(*s3.ListMultipartUploadsInput) (*request.Request, *s3.ListMultipartUploadsOutput) {
	return nil, nil
}

func (m mockS3Repository) ListMultipartUploadsPages(*s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool) error {
	return nil
}

func (m mockS3Repository) ListMultipartUploadsPagesWithContext(aws.Context, *s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m mockS3Repository) ListObjectVersions(*s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListObjectVersionsWithContext(aws.Context, *s3.ListObjectVersionsInput, ...request.Option) (*s3.ListObjectVersionsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListObjectVersionsRequest(*s3.ListObjectVersionsInput) (*request.Request, *s3.ListObjectVersionsOutput) {
	return nil, nil
}

func (m mockS3Repository) ListObjectVersionsPages(*s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool) error {
	return nil
}

func (m mockS3Repository) ListObjectVersionsPagesWithContext(aws.Context, *s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m mockS3Repository) ListObjects(*s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListObjectsWithContext(aws.Context, *s3.ListObjectsInput, ...request.Option) (*s3.ListObjectsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListObjectsRequest(*s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput) {
	return nil, nil
}

func (m mockS3Repository) ListObjectsPages(*s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool) error {
	return nil
}

func (m mockS3Repository) ListObjectsPagesWithContext(aws.Context, *s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m mockS3Repository) ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return nil, nil
}

func (m mockS3Repository) ListObjectsV2WithContext(aws.Context, *s3.ListObjectsV2Input, ...request.Option) (*s3.ListObjectsV2Output, error) {
	return nil, nil
}

func (m mockS3Repository) ListObjectsV2Request(*s3.ListObjectsV2Input) (*request.Request, *s3.ListObjectsV2Output) {
	return nil, nil
}

func (m mockS3Repository) ListObjectsV2Pages(*s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool) error {
	return nil
}

func (m mockS3Repository) ListObjectsV2PagesWithContext(aws.Context, *s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool, ...request.Option) error {
	return nil
}

func (m mockS3Repository) ListParts(*s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListPartsWithContext(aws.Context, *s3.ListPartsInput, ...request.Option) (*s3.ListPartsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) ListPartsRequest(*s3.ListPartsInput) (*request.Request, *s3.ListPartsOutput) {
	return nil, nil
}

func (m mockS3Repository) ListPartsPages(*s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool) error {
	return nil
}

func (m mockS3Repository) ListPartsPagesWithContext(aws.Context, *s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m mockS3Repository) PutBucketAccelerateConfiguration(*s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketAccelerateConfigurationWithContext(aws.Context, *s3.PutBucketAccelerateConfigurationInput, ...request.Option) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketAccelerateConfigurationRequest(*s3.PutBucketAccelerateConfigurationInput) (*request.Request, *s3.PutBucketAccelerateConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketAcl(*s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketAclWithContext(aws.Context, *s3.PutBucketAclInput, ...request.Option) (*s3.PutBucketAclOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketAclRequest(*s3.PutBucketAclInput) (*request.Request, *s3.PutBucketAclOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketAnalyticsConfiguration(*s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketAnalyticsConfigurationWithContext(aws.Context, *s3.PutBucketAnalyticsConfigurationInput, ...request.Option) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketAnalyticsConfigurationRequest(*s3.PutBucketAnalyticsConfigurationInput) (*request.Request, *s3.PutBucketAnalyticsConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketCors(*s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketCorsWithContext(aws.Context, *s3.PutBucketCorsInput, ...request.Option) (*s3.PutBucketCorsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketCorsRequest(*s3.PutBucketCorsInput) (*request.Request, *s3.PutBucketCorsOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketEncryption(*s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketEncryptionWithContext(aws.Context, *s3.PutBucketEncryptionInput, ...request.Option) (*s3.PutBucketEncryptionOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketEncryptionRequest(*s3.PutBucketEncryptionInput) (*request.Request, *s3.PutBucketEncryptionOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketIntelligentTieringConfiguration(*s3.PutBucketIntelligentTieringConfigurationInput) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketIntelligentTieringConfigurationWithContext(aws.Context, *s3.PutBucketIntelligentTieringConfigurationInput, ...request.Option) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketIntelligentTieringConfigurationRequest(*s3.PutBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.PutBucketIntelligentTieringConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketInventoryConfiguration(*s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketInventoryConfigurationWithContext(aws.Context, *s3.PutBucketInventoryConfigurationInput, ...request.Option) (*s3.PutBucketInventoryConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketInventoryConfigurationRequest(*s3.PutBucketInventoryConfigurationInput) (*request.Request, *s3.PutBucketInventoryConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketLifecycle(*s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketLifecycleWithContext(aws.Context, *s3.PutBucketLifecycleInput, ...request.Option) (*s3.PutBucketLifecycleOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketLifecycleRequest(*s3.PutBucketLifecycleInput) (*request.Request, *s3.PutBucketLifecycleOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketLifecycleConfiguration(*s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketLifecycleConfigurationWithContext(aws.Context, *s3.PutBucketLifecycleConfigurationInput, ...request.Option) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketLifecycleConfigurationRequest(*s3.PutBucketLifecycleConfigurationInput) (*request.Request, *s3.PutBucketLifecycleConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketLogging(*s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketLoggingWithContext(aws.Context, *s3.PutBucketLoggingInput, ...request.Option) (*s3.PutBucketLoggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketLoggingRequest(*s3.PutBucketLoggingInput) (*request.Request, *s3.PutBucketLoggingOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketMetricsConfiguration(*s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketMetricsConfigurationWithContext(aws.Context, *s3.PutBucketMetricsConfigurationInput, ...request.Option) (*s3.PutBucketMetricsConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketMetricsConfigurationRequest(*s3.PutBucketMetricsConfigurationInput) (*request.Request, *s3.PutBucketMetricsConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketNotification(*s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketNotificationWithContext(aws.Context, *s3.PutBucketNotificationInput, ...request.Option) (*s3.PutBucketNotificationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketNotificationRequest(*s3.PutBucketNotificationInput) (*request.Request, *s3.PutBucketNotificationOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketNotificationConfiguration(*s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketNotificationConfigurationWithContext(aws.Context, *s3.PutBucketNotificationConfigurationInput, ...request.Option) (*s3.PutBucketNotificationConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketNotificationConfigurationRequest(*s3.PutBucketNotificationConfigurationInput) (*request.Request, *s3.PutBucketNotificationConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketOwnershipControls(*s3.PutBucketOwnershipControlsInput) (*s3.PutBucketOwnershipControlsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketOwnershipControlsWithContext(aws.Context, *s3.PutBucketOwnershipControlsInput, ...request.Option) (*s3.PutBucketOwnershipControlsOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketOwnershipControlsRequest(*s3.PutBucketOwnershipControlsInput) (*request.Request, *s3.PutBucketOwnershipControlsOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketPolicy(*s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketPolicyWithContext(aws.Context, *s3.PutBucketPolicyInput, ...request.Option) (*s3.PutBucketPolicyOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketPolicyRequest(*s3.PutBucketPolicyInput) (*request.Request, *s3.PutBucketPolicyOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketReplication(*s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketReplicationWithContext(aws.Context, *s3.PutBucketReplicationInput, ...request.Option) (*s3.PutBucketReplicationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketReplicationRequest(*s3.PutBucketReplicationInput) (*request.Request, *s3.PutBucketReplicationOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketRequestPayment(*s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketRequestPaymentWithContext(aws.Context, *s3.PutBucketRequestPaymentInput, ...request.Option) (*s3.PutBucketRequestPaymentOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketRequestPaymentRequest(*s3.PutBucketRequestPaymentInput) (*request.Request, *s3.PutBucketRequestPaymentOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketTagging(*s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketTaggingWithContext(aws.Context, *s3.PutBucketTaggingInput, ...request.Option) (*s3.PutBucketTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketTaggingRequest(*s3.PutBucketTaggingInput) (*request.Request, *s3.PutBucketTaggingOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketVersioning(*s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketVersioningWithContext(aws.Context, *s3.PutBucketVersioningInput, ...request.Option) (*s3.PutBucketVersioningOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketVersioningRequest(*s3.PutBucketVersioningInput) (*request.Request, *s3.PutBucketVersioningOutput) {
	return nil, nil
}

func (m mockS3Repository) PutBucketWebsite(*s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketWebsiteWithContext(aws.Context, *s3.PutBucketWebsiteInput, ...request.Option) (*s3.PutBucketWebsiteOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutBucketWebsiteRequest(*s3.PutBucketWebsiteInput) (*request.Request, *s3.PutBucketWebsiteOutput) {
	return nil, nil
}

func (m mockS3Repository) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	if *input.Bucket != m.bucket {
		return nil, errorutil.Errorf(errorutil.Unknown, "bucket name is not valid")
	}

	if *input.Key != m.key {
		return nil, errorutil.Errorf(errorutil.Unknown, "key name is not valid")
	}

	return &s3.PutObjectOutput{}, nil
}

func (m mockS3Repository) PutObjectWithContext(aws.Context, *s3.PutObjectInput, ...request.Option) (*s3.PutObjectOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutObjectRequest(*s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	return nil, nil
}

func (m mockS3Repository) PutObjectAcl(*s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutObjectAclWithContext(aws.Context, *s3.PutObjectAclInput, ...request.Option) (*s3.PutObjectAclOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutObjectAclRequest(*s3.PutObjectAclInput) (*request.Request, *s3.PutObjectAclOutput) {
	return nil, nil
}

func (m mockS3Repository) PutObjectLegalHold(*s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutObjectLegalHoldWithContext(aws.Context, *s3.PutObjectLegalHoldInput, ...request.Option) (*s3.PutObjectLegalHoldOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutObjectLegalHoldRequest(*s3.PutObjectLegalHoldInput) (*request.Request, *s3.PutObjectLegalHoldOutput) {
	return nil, nil
}

func (m mockS3Repository) PutObjectLockConfiguration(*s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutObjectLockConfigurationWithContext(aws.Context, *s3.PutObjectLockConfigurationInput, ...request.Option) (*s3.PutObjectLockConfigurationOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutObjectLockConfigurationRequest(*s3.PutObjectLockConfigurationInput) (*request.Request, *s3.PutObjectLockConfigurationOutput) {
	return nil, nil
}

func (m mockS3Repository) PutObjectRetention(*s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutObjectRetentionWithContext(aws.Context, *s3.PutObjectRetentionInput, ...request.Option) (*s3.PutObjectRetentionOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutObjectRetentionRequest(*s3.PutObjectRetentionInput) (*request.Request, *s3.PutObjectRetentionOutput) {
	return nil, nil
}

func (m mockS3Repository) PutObjectTagging(*s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutObjectTaggingWithContext(aws.Context, *s3.PutObjectTaggingInput, ...request.Option) (*s3.PutObjectTaggingOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutObjectTaggingRequest(*s3.PutObjectTaggingInput) (*request.Request, *s3.PutObjectTaggingOutput) {
	return nil, nil
}

func (m mockS3Repository) PutPublicAccessBlock(*s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutPublicAccessBlockWithContext(aws.Context, *s3.PutPublicAccessBlockInput, ...request.Option) (*s3.PutPublicAccessBlockOutput, error) {
	return nil, nil
}

func (m mockS3Repository) PutPublicAccessBlockRequest(*s3.PutPublicAccessBlockInput) (*request.Request, *s3.PutPublicAccessBlockOutput) {
	return nil, nil
}

func (m mockS3Repository) RestoreObject(*s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	return nil, nil
}

func (m mockS3Repository) RestoreObjectWithContext(aws.Context, *s3.RestoreObjectInput, ...request.Option) (*s3.RestoreObjectOutput, error) {
	return nil, nil
}

func (m mockS3Repository) RestoreObjectRequest(*s3.RestoreObjectInput) (*request.Request, *s3.RestoreObjectOutput) {
	return nil, nil
}

func (m mockS3Repository) SelectObjectContent(*s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error) {
	return nil, nil
}

func (m mockS3Repository) SelectObjectContentWithContext(aws.Context, *s3.SelectObjectContentInput, ...request.Option) (*s3.SelectObjectContentOutput, error) {
	return nil, nil
}

func (m mockS3Repository) SelectObjectContentRequest(*s3.SelectObjectContentInput) (*request.Request, *s3.SelectObjectContentOutput) {
	return nil, nil
}

func (m mockS3Repository) UploadPart(*s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	return nil, nil
}

func (m mockS3Repository) UploadPartWithContext(aws.Context, *s3.UploadPartInput, ...request.Option) (*s3.UploadPartOutput, error) {
	return nil, nil
}

func (m mockS3Repository) UploadPartRequest(*s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput) {
	return nil, nil
}

func (m mockS3Repository) UploadPartCopy(*s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	return nil, nil
}

func (m mockS3Repository) UploadPartCopyWithContext(aws.Context, *s3.UploadPartCopyInput, ...request.Option) (*s3.UploadPartCopyOutput, error) {
	return nil, nil
}

func (m mockS3Repository) UploadPartCopyRequest(*s3.UploadPartCopyInput) (*request.Request, *s3.UploadPartCopyOutput) {
	return nil, nil
}

func (m mockS3Repository) WriteGetObjectResponse(*s3.WriteGetObjectResponseInput) (*s3.WriteGetObjectResponseOutput, error) {
	return nil, nil
}

func (m mockS3Repository) WriteGetObjectResponseWithContext(aws.Context, *s3.WriteGetObjectResponseInput, ...request.Option) (*s3.WriteGetObjectResponseOutput, error) {
	return nil, nil
}

func (m mockS3Repository) WriteGetObjectResponseRequest(*s3.WriteGetObjectResponseInput) (*request.Request, *s3.WriteGetObjectResponseOutput) {
	return nil, nil
}

func (m mockS3Repository) WaitUntilBucketExists(*s3.HeadBucketInput) error {
	return nil
}

func (m mockS3Repository) WaitUntilBucketExistsWithContext(aws.Context, *s3.HeadBucketInput, ...request.WaiterOption) error {
	return nil
}

func (m mockS3Repository) WaitUntilBucketNotExists(*s3.HeadBucketInput) error {
	return nil
}

func (m mockS3Repository) WaitUntilBucketNotExistsWithContext(aws.Context, *s3.HeadBucketInput, ...request.WaiterOption) error {
	return nil
}

func (m mockS3Repository) WaitUntilObjectExists(*s3.HeadObjectInput) error {
	return nil
}

func (m mockS3Repository) WaitUntilObjectExistsWithContext(aws.Context, *s3.HeadObjectInput, ...request.WaiterOption) error {
	return nil
}

func (m mockS3Repository) WaitUntilObjectNotExists(*s3.HeadObjectInput) error {
	return nil
}

func (m mockS3Repository) WaitUntilObjectNotExistsWithContext(aws.Context, *s3.HeadObjectInput, ...request.WaiterOption) error {
	return nil
}
