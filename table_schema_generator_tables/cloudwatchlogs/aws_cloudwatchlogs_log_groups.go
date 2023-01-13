package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCloudwatchlogsLogGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudwatchlogsLogGroupsGenerator{}

func (x *TableAwsCloudwatchlogsLogGroupsGenerator) GetTableName() string {
	return "aws_cloudwatchlogs_log_groups"
}

func (x *TableAwsCloudwatchlogsLogGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudwatchlogsLogGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudwatchlogsLogGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsCloudwatchlogsLogGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config cloudwatchlogs.DescribeLogGroupsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudwatchlogs
			for {
				response, err := svc.DescribeLogGroups(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.LogGroups
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsCloudwatchlogsLogGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("logs")
}

func (x *TableAwsCloudwatchlogsLogGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LogGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_in_days").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RetentionInDays")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_protection_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DataProtectionStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_filter_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MetricFilterCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stored_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("StoredBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
	}
}

func (x *TableAwsCloudwatchlogsLogGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
