package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsFrauddetectorBatchPredictionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsFrauddetectorBatchPredictionsGenerator{}

func (x *TableAwsFrauddetectorBatchPredictionsGenerator) GetTableName() string {
	return "aws_frauddetector_batch_predictions"
}

func (x *TableAwsFrauddetectorBatchPredictionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsFrauddetectorBatchPredictionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsFrauddetectorBatchPredictionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsFrauddetectorBatchPredictionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := frauddetector.NewGetBatchPredictionJobsPaginator(client.(*aws_client.Client).AwsServices().Frauddetector, nil)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.BatchPredictions
			}
			return nil
		},
	}
}

func (x *TableAwsFrauddetectorBatchPredictionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("frauddetector")
}

func (x *TableAwsFrauddetectorBatchPredictionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("detector_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DetectorVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("detector_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DetectorName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IamRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("processed_records_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ProcessedRecordsCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_type_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventTypeName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failure_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FailureReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("input_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InputPath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("job_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("JobId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_records_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TotalRecordsCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("completion_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CompletionTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_heartbeat_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastHeartbeatTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OutputPath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
	}
}

func (x *TableAwsFrauddetectorBatchPredictionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
