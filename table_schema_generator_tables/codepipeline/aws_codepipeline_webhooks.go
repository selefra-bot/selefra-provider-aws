package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCodepipelineWebhooksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCodepipelineWebhooksGenerator{}

func (x *TableAwsCodepipelineWebhooksGenerator) GetTableName() string {
	return "aws_codepipeline_webhooks"
}

func (x *TableAwsCodepipelineWebhooksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCodepipelineWebhooksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCodepipelineWebhooksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsCodepipelineWebhooksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Codepipeline
			config := codepipeline.ListWebhooksInput{}
			for {
				response, err := svc.ListWebhooks(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Webhooks

				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsCodepipelineWebhooksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("codepipeline")
}

func (x *TableAwsCodepipelineWebhooksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("definition").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Definition")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_triggered").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastTriggered")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Url")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ErrorCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ErrorMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsCodepipelineWebhooksGenerator) GetSubTables() []*schema.Table {
	return nil
}
