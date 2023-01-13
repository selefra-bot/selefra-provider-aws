package inspector2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsInspector2FindingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsInspector2FindingsGenerator{}

func (x *TableAwsInspector2FindingsGenerator) GetTableName() string {
	return "aws_inspector2_findings"
}

func (x *TableAwsInspector2FindingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsInspector2FindingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsInspector2FindingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsInspector2FindingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Inspector2
			input := inspector2.ListFindingsInput{MaxResults: aws.Int32(100)}
			for {
				response, err := svc.ListFindings(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Findings
				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsInspector2FindingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("inspector2")
}

func (x *TableAwsInspector2FindingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("aws_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AwsAccountId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inspector_score").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("InspectorScore")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("finding_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FindingArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("remediation").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Remediation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Resources")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inspector_score_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("InspectorScoreDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FindingArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fix_available").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FixAvailable")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_reachability_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NetworkReachabilityDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("first_observed_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("FirstObservedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_observed_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastObservedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("severity").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Severity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("exploit_available").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExploitAvailable")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("exploitability_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExploitabilityDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("package_vulnerability_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PackageVulnerabilityDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Title")).Build(),
	}
}

func (x *TableAwsInspector2FindingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
