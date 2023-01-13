package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsInspectorFindingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsInspectorFindingsGenerator{}

func (x *TableAwsInspectorFindingsGenerator) GetTableName() string {
	return "aws_inspector_findings"
}

func (x *TableAwsInspectorFindingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsInspectorFindingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsInspectorFindingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsInspectorFindingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Inspector
			input := inspector.ListFindingsInput{MaxResults: aws.Int32(100)}
			for {
				response, err := svc.ListFindings(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				if len(response.FindingArns) > 0 {
					out, err := svc.DescribeFindings(ctx, &inspector.DescribeFindingsInput{FindingArns: response.FindingArns})
					if err != nil {
						if c.IsNotFoundError(err) {
							continue
						}
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- out.Findings
				}
				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsInspectorFindingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("inspector")
}

func (x *TableAwsInspectorFindingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("confidence").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Confidence")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_version").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SchemaVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Title")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("severity").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Severity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("indicator_of_compromise").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IndicatorOfCompromise")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recommendation").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Recommendation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Attributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UserAttributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("asset_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AssetType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServiceAttributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("asset_attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AssetAttributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("numeric_severity").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("NumericSeverity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Service")).Build(),
	}
}

func (x *TableAwsInspectorFindingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
