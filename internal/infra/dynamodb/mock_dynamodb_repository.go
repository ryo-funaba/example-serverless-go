package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	dynamoDB "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
)

type mockDynamodbRepository struct {
	Table string
	PK    string
	SK    string
}

func (m mockDynamodbRepository) BatchExecuteStatement(*dynamoDB.BatchExecuteStatementInput) (*dynamoDB.BatchExecuteStatementOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) BatchExecuteStatementWithContext(aws.Context, *dynamoDB.BatchExecuteStatementInput, ...request.Option) (*dynamoDB.BatchExecuteStatementOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) BatchExecuteStatementRequest(*dynamoDB.BatchExecuteStatementInput) (*request.Request, *dynamoDB.BatchExecuteStatementOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) BatchGetItem(*dynamoDB.BatchGetItemInput) (*dynamoDB.BatchGetItemOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) BatchGetItemWithContext(aws.Context, *dynamoDB.BatchGetItemInput, ...request.Option) (*dynamoDB.BatchGetItemOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) BatchGetItemRequest(*dynamoDB.BatchGetItemInput) (*request.Request, *dynamoDB.BatchGetItemOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) BatchGetItemPages(*dynamoDB.BatchGetItemInput, func(*dynamoDB.BatchGetItemOutput, bool) bool) error {
	return nil
}

func (m mockDynamodbRepository) BatchGetItemPagesWithContext(aws.Context, *dynamoDB.BatchGetItemInput, func(*dynamoDB.BatchGetItemOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m mockDynamodbRepository) BatchWriteItem(*dynamoDB.BatchWriteItemInput) (*dynamoDB.BatchWriteItemOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) BatchWriteItemWithContext(aws.Context, *dynamoDB.BatchWriteItemInput, ...request.Option) (*dynamoDB.BatchWriteItemOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) BatchWriteItemRequest(*dynamoDB.BatchWriteItemInput) (*request.Request, *dynamoDB.BatchWriteItemOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) CreateBackup(*dynamoDB.CreateBackupInput) (*dynamoDB.CreateBackupOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) CreateBackupWithContext(aws.Context, *dynamoDB.CreateBackupInput, ...request.Option) (*dynamoDB.CreateBackupOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) CreateBackupRequest(*dynamoDB.CreateBackupInput) (*request.Request, *dynamoDB.CreateBackupOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) CreateGlobalTable(*dynamoDB.CreateGlobalTableInput) (*dynamoDB.CreateGlobalTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) CreateGlobalTableWithContext(aws.Context, *dynamoDB.CreateGlobalTableInput, ...request.Option) (*dynamoDB.CreateGlobalTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) CreateGlobalTableRequest(*dynamoDB.CreateGlobalTableInput) (*request.Request, *dynamoDB.CreateGlobalTableOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) CreateTable(*dynamoDB.CreateTableInput) (*dynamoDB.CreateTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) CreateTableWithContext(aws.Context, *dynamoDB.CreateTableInput, ...request.Option) (*dynamoDB.CreateTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) CreateTableRequest(*dynamoDB.CreateTableInput) (*request.Request, *dynamoDB.CreateTableOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DeleteBackup(*dynamoDB.DeleteBackupInput) (*dynamoDB.DeleteBackupOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DeleteBackupWithContext(aws.Context, *dynamoDB.DeleteBackupInput, ...request.Option) (*dynamoDB.DeleteBackupOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DeleteBackupRequest(*dynamoDB.DeleteBackupInput) (*request.Request, *dynamoDB.DeleteBackupOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DeleteItem(*dynamoDB.DeleteItemInput) (*dynamoDB.DeleteItemOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DeleteItemWithContext(aws.Context, *dynamoDB.DeleteItemInput, ...request.Option) (*dynamoDB.DeleteItemOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DeleteItemRequest(*dynamoDB.DeleteItemInput) (*request.Request, *dynamoDB.DeleteItemOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DeleteTable(*dynamoDB.DeleteTableInput) (*dynamoDB.DeleteTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DeleteTableWithContext(aws.Context, *dynamoDB.DeleteTableInput, ...request.Option) (*dynamoDB.DeleteTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DeleteTableRequest(*dynamoDB.DeleteTableInput) (*request.Request, *dynamoDB.DeleteTableOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeBackup(*dynamoDB.DescribeBackupInput) (*dynamoDB.DescribeBackupOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeBackupWithContext(aws.Context, *dynamoDB.DescribeBackupInput, ...request.Option) (*dynamoDB.DescribeBackupOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeBackupRequest(*dynamoDB.DescribeBackupInput) (*request.Request, *dynamoDB.DescribeBackupOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeContinuousBackups(*dynamoDB.DescribeContinuousBackupsInput) (*dynamoDB.DescribeContinuousBackupsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeContinuousBackupsWithContext(aws.Context, *dynamoDB.DescribeContinuousBackupsInput, ...request.Option) (*dynamoDB.DescribeContinuousBackupsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeContinuousBackupsRequest(*dynamoDB.DescribeContinuousBackupsInput) (*request.Request, *dynamoDB.DescribeContinuousBackupsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeContributorInsights(*dynamoDB.DescribeContributorInsightsInput) (*dynamoDB.DescribeContributorInsightsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeContributorInsightsWithContext(aws.Context, *dynamoDB.DescribeContributorInsightsInput, ...request.Option) (*dynamoDB.DescribeContributorInsightsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeContributorInsightsRequest(*dynamoDB.DescribeContributorInsightsInput) (*request.Request, *dynamoDB.DescribeContributorInsightsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeEndpoints(*dynamoDB.DescribeEndpointsInput) (*dynamoDB.DescribeEndpointsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeEndpointsWithContext(aws.Context, *dynamoDB.DescribeEndpointsInput, ...request.Option) (*dynamoDB.DescribeEndpointsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeEndpointsRequest(*dynamoDB.DescribeEndpointsInput) (*request.Request, *dynamoDB.DescribeEndpointsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeExport(*dynamoDB.DescribeExportInput) (*dynamoDB.DescribeExportOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeExportWithContext(aws.Context, *dynamoDB.DescribeExportInput, ...request.Option) (*dynamoDB.DescribeExportOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeExportRequest(*dynamoDB.DescribeExportInput) (*request.Request, *dynamoDB.DescribeExportOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeGlobalTable(*dynamoDB.DescribeGlobalTableInput) (*dynamoDB.DescribeGlobalTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeGlobalTableWithContext(aws.Context, *dynamoDB.DescribeGlobalTableInput, ...request.Option) (*dynamoDB.DescribeGlobalTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeGlobalTableRequest(*dynamoDB.DescribeGlobalTableInput) (*request.Request, *dynamoDB.DescribeGlobalTableOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeGlobalTableSettings(*dynamoDB.DescribeGlobalTableSettingsInput) (*dynamoDB.DescribeGlobalTableSettingsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeGlobalTableSettingsWithContext(aws.Context, *dynamoDB.DescribeGlobalTableSettingsInput, ...request.Option) (*dynamoDB.DescribeGlobalTableSettingsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeGlobalTableSettingsRequest(*dynamoDB.DescribeGlobalTableSettingsInput) (*request.Request, *dynamoDB.DescribeGlobalTableSettingsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeImport(*dynamoDB.DescribeImportInput) (*dynamoDB.DescribeImportOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeImportWithContext(aws.Context, *dynamoDB.DescribeImportInput, ...request.Option) (*dynamoDB.DescribeImportOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeImportRequest(*dynamoDB.DescribeImportInput) (*request.Request, *dynamoDB.DescribeImportOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeKinesisStreamingDestination(*dynamoDB.DescribeKinesisStreamingDestinationInput) (*dynamoDB.DescribeKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeKinesisStreamingDestinationWithContext(aws.Context, *dynamoDB.DescribeKinesisStreamingDestinationInput, ...request.Option) (*dynamoDB.DescribeKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeKinesisStreamingDestinationRequest(*dynamoDB.DescribeKinesisStreamingDestinationInput) (*request.Request, *dynamoDB.DescribeKinesisStreamingDestinationOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeLimits(*dynamoDB.DescribeLimitsInput) (*dynamoDB.DescribeLimitsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeLimitsWithContext(aws.Context, *dynamoDB.DescribeLimitsInput, ...request.Option) (*dynamoDB.DescribeLimitsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeLimitsRequest(*dynamoDB.DescribeLimitsInput) (*request.Request, *dynamoDB.DescribeLimitsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeTable(*dynamoDB.DescribeTableInput) (*dynamoDB.DescribeTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeTableWithContext(aws.Context, *dynamoDB.DescribeTableInput, ...request.Option) (*dynamoDB.DescribeTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeTableRequest(*dynamoDB.DescribeTableInput) (*request.Request, *dynamoDB.DescribeTableOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeTableReplicaAutoScaling(*dynamoDB.DescribeTableReplicaAutoScalingInput) (*dynamoDB.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeTableReplicaAutoScalingWithContext(aws.Context, *dynamoDB.DescribeTableReplicaAutoScalingInput, ...request.Option) (*dynamoDB.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeTableReplicaAutoScalingRequest(*dynamoDB.DescribeTableReplicaAutoScalingInput) (*request.Request, *dynamoDB.DescribeTableReplicaAutoScalingOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeTimeToLive(*dynamoDB.DescribeTimeToLiveInput) (*dynamoDB.DescribeTimeToLiveOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeTimeToLiveWithContext(aws.Context, *dynamoDB.DescribeTimeToLiveInput, ...request.Option) (*dynamoDB.DescribeTimeToLiveOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DescribeTimeToLiveRequest(*dynamoDB.DescribeTimeToLiveInput) (*request.Request, *dynamoDB.DescribeTimeToLiveOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) DisableKinesisStreamingDestination(*dynamoDB.DisableKinesisStreamingDestinationInput) (*dynamoDB.DisableKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DisableKinesisStreamingDestinationWithContext(aws.Context, *dynamoDB.DisableKinesisStreamingDestinationInput, ...request.Option) (*dynamoDB.DisableKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) DisableKinesisStreamingDestinationRequest(*dynamoDB.DisableKinesisStreamingDestinationInput) (*request.Request, *dynamoDB.DisableKinesisStreamingDestinationOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) EnableKinesisStreamingDestination(*dynamoDB.EnableKinesisStreamingDestinationInput) (*dynamoDB.EnableKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) EnableKinesisStreamingDestinationWithContext(aws.Context, *dynamoDB.EnableKinesisStreamingDestinationInput, ...request.Option) (*dynamoDB.EnableKinesisStreamingDestinationOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) EnableKinesisStreamingDestinationRequest(*dynamoDB.EnableKinesisStreamingDestinationInput) (*request.Request, *dynamoDB.EnableKinesisStreamingDestinationOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ExecuteStatement(*dynamoDB.ExecuteStatementInput) (*dynamoDB.ExecuteStatementOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ExecuteStatementWithContext(aws.Context, *dynamoDB.ExecuteStatementInput, ...request.Option) (*dynamoDB.ExecuteStatementOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ExecuteStatementRequest(*dynamoDB.ExecuteStatementInput) (*request.Request, *dynamoDB.ExecuteStatementOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ExecuteTransaction(*dynamoDB.ExecuteTransactionInput) (*dynamoDB.ExecuteTransactionOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ExecuteTransactionWithContext(aws.Context, *dynamoDB.ExecuteTransactionInput, ...request.Option) (*dynamoDB.ExecuteTransactionOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ExecuteTransactionRequest(*dynamoDB.ExecuteTransactionInput) (*request.Request, *dynamoDB.ExecuteTransactionOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ExportTableToPointInTime(*dynamoDB.ExportTableToPointInTimeInput) (*dynamoDB.ExportTableToPointInTimeOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ExportTableToPointInTimeWithContext(aws.Context, *dynamoDB.ExportTableToPointInTimeInput, ...request.Option) (*dynamoDB.ExportTableToPointInTimeOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ExportTableToPointInTimeRequest(*dynamoDB.ExportTableToPointInTimeInput) (*request.Request, *dynamoDB.ExportTableToPointInTimeOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) GetItem(input *dynamoDB.GetItemInput) (*dynamoDB.GetItemOutput, error) {
	if m.Table != *input.TableName {
		return nil, errorutil.Errorf(errorutil.Unknown, "TableName is invalid")
	}

	if m.PK != *input.Key["PK"].S {
		return nil, errorutil.Errorf(errorutil.Unknown, "PK is invalid")
	}

	if m.SK != *input.Key["SK"].S {
		return nil, errorutil.Errorf(errorutil.Unknown, "SK is invalid")
	}

	return &dynamoDB.GetItemOutput{}, nil
}

func (m mockDynamodbRepository) GetItemWithContext(aws.Context, *dynamoDB.GetItemInput, ...request.Option) (*dynamoDB.GetItemOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) GetItemRequest(*dynamoDB.GetItemInput) (*request.Request, *dynamoDB.GetItemOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ImportTable(*dynamoDB.ImportTableInput) (*dynamoDB.ImportTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ImportTableWithContext(aws.Context, *dynamoDB.ImportTableInput, ...request.Option) (*dynamoDB.ImportTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ImportTableRequest(*dynamoDB.ImportTableInput) (*request.Request, *dynamoDB.ImportTableOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ListBackups(*dynamoDB.ListBackupsInput) (*dynamoDB.ListBackupsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListBackupsWithContext(aws.Context, *dynamoDB.ListBackupsInput, ...request.Option) (*dynamoDB.ListBackupsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListBackupsRequest(*dynamoDB.ListBackupsInput) (*request.Request, *dynamoDB.ListBackupsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ListContributorInsights(*dynamoDB.ListContributorInsightsInput) (*dynamoDB.ListContributorInsightsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListContributorInsightsWithContext(aws.Context, *dynamoDB.ListContributorInsightsInput, ...request.Option) (*dynamoDB.ListContributorInsightsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListContributorInsightsRequest(*dynamoDB.ListContributorInsightsInput) (*request.Request, *dynamoDB.ListContributorInsightsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ListContributorInsightsPages(*dynamoDB.ListContributorInsightsInput, func(*dynamoDB.ListContributorInsightsOutput, bool) bool) error {
	return nil
}

func (m mockDynamodbRepository) ListContributorInsightsPagesWithContext(aws.Context, *dynamoDB.ListContributorInsightsInput, func(*dynamoDB.ListContributorInsightsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m mockDynamodbRepository) ListExports(*dynamoDB.ListExportsInput) (*dynamoDB.ListExportsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListExportsWithContext(aws.Context, *dynamoDB.ListExportsInput, ...request.Option) (*dynamoDB.ListExportsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListExportsRequest(*dynamoDB.ListExportsInput) (*request.Request, *dynamoDB.ListExportsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ListExportsPages(*dynamoDB.ListExportsInput, func(*dynamoDB.ListExportsOutput, bool) bool) error {
	return nil
}

func (m mockDynamodbRepository) ListExportsPagesWithContext(aws.Context, *dynamoDB.ListExportsInput, func(*dynamoDB.ListExportsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m mockDynamodbRepository) ListGlobalTables(*dynamoDB.ListGlobalTablesInput) (*dynamoDB.ListGlobalTablesOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListGlobalTablesWithContext(aws.Context, *dynamoDB.ListGlobalTablesInput, ...request.Option) (*dynamoDB.ListGlobalTablesOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListGlobalTablesRequest(*dynamoDB.ListGlobalTablesInput) (*request.Request, *dynamoDB.ListGlobalTablesOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ListImports(*dynamoDB.ListImportsInput) (*dynamoDB.ListImportsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListImportsWithContext(aws.Context, *dynamoDB.ListImportsInput, ...request.Option) (*dynamoDB.ListImportsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListImportsRequest(*dynamoDB.ListImportsInput) (*request.Request, *dynamoDB.ListImportsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ListImportsPages(*dynamoDB.ListImportsInput, func(*dynamoDB.ListImportsOutput, bool) bool) error {
	return nil
}

func (m mockDynamodbRepository) ListImportsPagesWithContext(aws.Context, *dynamoDB.ListImportsInput, func(*dynamoDB.ListImportsOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m mockDynamodbRepository) ListTables(*dynamoDB.ListTablesInput) (*dynamoDB.ListTablesOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListTablesWithContext(aws.Context, *dynamoDB.ListTablesInput, ...request.Option) (*dynamoDB.ListTablesOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListTablesRequest(*dynamoDB.ListTablesInput) (*request.Request, *dynamoDB.ListTablesOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ListTablesPages(*dynamoDB.ListTablesInput, func(*dynamoDB.ListTablesOutput, bool) bool) error {
	return nil
}

func (m mockDynamodbRepository) ListTablesPagesWithContext(aws.Context, *dynamoDB.ListTablesInput, func(*dynamoDB.ListTablesOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m mockDynamodbRepository) ListTagsOfResource(*dynamoDB.ListTagsOfResourceInput) (*dynamoDB.ListTagsOfResourceOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListTagsOfResourceWithContext(aws.Context, *dynamoDB.ListTagsOfResourceInput, ...request.Option) (*dynamoDB.ListTagsOfResourceOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ListTagsOfResourceRequest(*dynamoDB.ListTagsOfResourceInput) (*request.Request, *dynamoDB.ListTagsOfResourceOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) PutItem(input *dynamoDB.PutItemInput) (*dynamoDB.PutItemOutput, error) {
	if m.Table != *input.TableName {
		return nil, errorutil.Errorf(errorutil.Unknown, "TableName is invalid")
	}

	if m.PK != *input.Item["PK"].S {
		return nil, errorutil.Errorf(errorutil.Unknown, "PK is invalid")
	}

	if m.SK != *input.Item["SK"].S {
		return nil, errorutil.Errorf(errorutil.Unknown, "SK is invalid")
	}

	return &dynamoDB.PutItemOutput{}, nil
}

func (m mockDynamodbRepository) PutItemWithContext(aws.Context, *dynamoDB.PutItemInput, ...request.Option) (*dynamoDB.PutItemOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) PutItemRequest(*dynamoDB.PutItemInput) (*request.Request, *dynamoDB.PutItemOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) Query(*dynamoDB.QueryInput) (*dynamoDB.QueryOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) QueryWithContext(aws.Context, *dynamoDB.QueryInput, ...request.Option) (*dynamoDB.QueryOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) QueryRequest(*dynamoDB.QueryInput) (*request.Request, *dynamoDB.QueryOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) QueryPages(*dynamoDB.QueryInput, func(*dynamoDB.QueryOutput, bool) bool) error {
	return nil
}

func (m mockDynamodbRepository) QueryPagesWithContext(aws.Context, *dynamoDB.QueryInput, func(*dynamoDB.QueryOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m mockDynamodbRepository) RestoreTableFromBackup(*dynamoDB.RestoreTableFromBackupInput) (*dynamoDB.RestoreTableFromBackupOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) RestoreTableFromBackupWithContext(aws.Context, *dynamoDB.RestoreTableFromBackupInput, ...request.Option) (*dynamoDB.RestoreTableFromBackupOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) RestoreTableFromBackupRequest(*dynamoDB.RestoreTableFromBackupInput) (*request.Request, *dynamoDB.RestoreTableFromBackupOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) RestoreTableToPointInTime(*dynamoDB.RestoreTableToPointInTimeInput) (*dynamoDB.RestoreTableToPointInTimeOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) RestoreTableToPointInTimeWithContext(aws.Context, *dynamoDB.RestoreTableToPointInTimeInput, ...request.Option) (*dynamoDB.RestoreTableToPointInTimeOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) RestoreTableToPointInTimeRequest(*dynamoDB.RestoreTableToPointInTimeInput) (*request.Request, *dynamoDB.RestoreTableToPointInTimeOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) Scan(*dynamoDB.ScanInput) (*dynamoDB.ScanOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ScanWithContext(aws.Context, *dynamoDB.ScanInput, ...request.Option) (*dynamoDB.ScanOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) ScanRequest(*dynamoDB.ScanInput) (*request.Request, *dynamoDB.ScanOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) ScanPages(*dynamoDB.ScanInput, func(*dynamoDB.ScanOutput, bool) bool) error {
	return nil
}

func (m mockDynamodbRepository) ScanPagesWithContext(aws.Context, *dynamoDB.ScanInput, func(*dynamoDB.ScanOutput, bool) bool, ...request.Option) error {
	return nil
}

func (m mockDynamodbRepository) TagResource(*dynamoDB.TagResourceInput) (*dynamoDB.TagResourceOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) TagResourceWithContext(aws.Context, *dynamoDB.TagResourceInput, ...request.Option) (*dynamoDB.TagResourceOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) TagResourceRequest(*dynamoDB.TagResourceInput) (*request.Request, *dynamoDB.TagResourceOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) TransactGetItems(*dynamoDB.TransactGetItemsInput) (*dynamoDB.TransactGetItemsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) TransactGetItemsWithContext(aws.Context, *dynamoDB.TransactGetItemsInput, ...request.Option) (*dynamoDB.TransactGetItemsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) TransactGetItemsRequest(*dynamoDB.TransactGetItemsInput) (*request.Request, *dynamoDB.TransactGetItemsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) TransactWriteItems(*dynamoDB.TransactWriteItemsInput) (*dynamoDB.TransactWriteItemsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) TransactWriteItemsWithContext(aws.Context, *dynamoDB.TransactWriteItemsInput, ...request.Option) (*dynamoDB.TransactWriteItemsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) TransactWriteItemsRequest(*dynamoDB.TransactWriteItemsInput) (*request.Request, *dynamoDB.TransactWriteItemsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) UntagResource(*dynamoDB.UntagResourceInput) (*dynamoDB.UntagResourceOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UntagResourceWithContext(aws.Context, *dynamoDB.UntagResourceInput, ...request.Option) (*dynamoDB.UntagResourceOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UntagResourceRequest(*dynamoDB.UntagResourceInput) (*request.Request, *dynamoDB.UntagResourceOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateContinuousBackups(*dynamoDB.UpdateContinuousBackupsInput) (*dynamoDB.UpdateContinuousBackupsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateContinuousBackupsWithContext(aws.Context, *dynamoDB.UpdateContinuousBackupsInput, ...request.Option) (*dynamoDB.UpdateContinuousBackupsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateContinuousBackupsRequest(*dynamoDB.UpdateContinuousBackupsInput) (*request.Request, *dynamoDB.UpdateContinuousBackupsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateContributorInsights(*dynamoDB.UpdateContributorInsightsInput) (*dynamoDB.UpdateContributorInsightsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateContributorInsightsWithContext(aws.Context, *dynamoDB.UpdateContributorInsightsInput, ...request.Option) (*dynamoDB.UpdateContributorInsightsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateContributorInsightsRequest(*dynamoDB.UpdateContributorInsightsInput) (*request.Request, *dynamoDB.UpdateContributorInsightsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateGlobalTable(*dynamoDB.UpdateGlobalTableInput) (*dynamoDB.UpdateGlobalTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateGlobalTableWithContext(aws.Context, *dynamoDB.UpdateGlobalTableInput, ...request.Option) (*dynamoDB.UpdateGlobalTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateGlobalTableRequest(*dynamoDB.UpdateGlobalTableInput) (*request.Request, *dynamoDB.UpdateGlobalTableOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateGlobalTableSettings(*dynamoDB.UpdateGlobalTableSettingsInput) (*dynamoDB.UpdateGlobalTableSettingsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateGlobalTableSettingsWithContext(aws.Context, *dynamoDB.UpdateGlobalTableSettingsInput, ...request.Option) (*dynamoDB.UpdateGlobalTableSettingsOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateGlobalTableSettingsRequest(*dynamoDB.UpdateGlobalTableSettingsInput) (*request.Request, *dynamoDB.UpdateGlobalTableSettingsOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateItem(*dynamoDB.UpdateItemInput) (*dynamoDB.UpdateItemOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateItemWithContext(aws.Context, *dynamoDB.UpdateItemInput, ...request.Option) (*dynamoDB.UpdateItemOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateItemRequest(*dynamoDB.UpdateItemInput) (*request.Request, *dynamoDB.UpdateItemOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateTable(*dynamoDB.UpdateTableInput) (*dynamoDB.UpdateTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateTableWithContext(aws.Context, *dynamoDB.UpdateTableInput, ...request.Option) (*dynamoDB.UpdateTableOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateTableRequest(*dynamoDB.UpdateTableInput) (*request.Request, *dynamoDB.UpdateTableOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateTableReplicaAutoScaling(*dynamoDB.UpdateTableReplicaAutoScalingInput) (*dynamoDB.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateTableReplicaAutoScalingWithContext(aws.Context, *dynamoDB.UpdateTableReplicaAutoScalingInput, ...request.Option) (*dynamoDB.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateTableReplicaAutoScalingRequest(*dynamoDB.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamoDB.UpdateTableReplicaAutoScalingOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateTimeToLive(*dynamoDB.UpdateTimeToLiveInput) (*dynamoDB.UpdateTimeToLiveOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateTimeToLiveWithContext(aws.Context, *dynamoDB.UpdateTimeToLiveInput, ...request.Option) (*dynamoDB.UpdateTimeToLiveOutput, error) {
	return nil, nil
}

func (m mockDynamodbRepository) UpdateTimeToLiveRequest(*dynamoDB.UpdateTimeToLiveInput) (*request.Request, *dynamoDB.UpdateTimeToLiveOutput) {
	return nil, nil
}

func (m mockDynamodbRepository) WaitUntilTableExists(*dynamoDB.DescribeTableInput) error {
	return nil
}

func (m mockDynamodbRepository) WaitUntilTableExistsWithContext(aws.Context, *dynamoDB.DescribeTableInput, ...request.WaiterOption) error {
	return nil
}

func (m mockDynamodbRepository) WaitUntilTableNotExists(*dynamoDB.DescribeTableInput) error {
	return nil
}

func (m mockDynamodbRepository) WaitUntilTableNotExistsWithContext(aws.Context, *dynamoDB.DescribeTableInput, ...request.WaiterOption) error {
	return nil
}
