package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/smithy-go"
	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsQuicksightDashboardsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsQuicksightDashboardsGenerator{}

func (x *TableAwsQuicksightDashboardsGenerator) GetTableName() string {
	return "aws_quicksight_dashboards"
}

func (x *TableAwsQuicksightDashboardsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsQuicksightDashboardsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsQuicksightDashboardsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsQuicksightDashboardsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Quicksight
			input := quicksight.ListDashboardsInput{
				AwsAccountId: aws.String(cl.AccountID),
			}
			var ae smithy.APIError

			paginator := quicksight.NewListDashboardsPaginator(svc, &input)
			for paginator.HasMorePages() {
				result, err := paginator.NextPage(ctx)
				if err != nil {
					if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.DashboardSummaryList
			}
			return nil
		},
	}
}

func (x *TableAwsQuicksightDashboardsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("quicksight")
}

func (x *TableAwsQuicksightDashboardsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("published_version_number").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("PublishedVersionNumber")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dashboard_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DashboardId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_published_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastPublishedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
	}
}

func (x *TableAwsQuicksightDashboardsGenerator) GetSubTables() []*schema.Table {
	return nil
}
