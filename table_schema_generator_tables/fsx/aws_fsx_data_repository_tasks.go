package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsFsxDataRepositoryTasksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsFsxDataRepositoryTasksGenerator{}

func (x *TableAwsFsxDataRepositoryTasksGenerator) GetTableName() string {
	return "aws_fsx_data_repository_tasks"
}

func (x *TableAwsFsxDataRepositoryTasksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsFsxDataRepositoryTasksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsFsxDataRepositoryTasksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsFsxDataRepositoryTasksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Fsx
			input := fsx.DescribeDataRepositoryTasksInput{MaxResults: aws.Int32(1000)}
			paginator := fsx.NewDescribeDataRepositoryTasksPaginator(svc, &input)
			for paginator.HasMorePages() {
				result, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.DataRepositoryTasks
			}
			return nil
		},
	}
}

func (x *TableAwsFsxDataRepositoryTasksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("fsx")
}

func (x *TableAwsFsxDataRepositoryTasksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("resource_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("EndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failure_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FailureDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("task_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TaskId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("paths").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Paths")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("StartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity_to_release").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("CapacityToRelease")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_cache_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FileCacheId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_system_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FileSystemId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("report").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Report")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Lifecycle")).Build(),
	}
}

func (x *TableAwsFsxDataRepositoryTasksGenerator) GetSubTables() []*schema.Table {
	return nil
}
