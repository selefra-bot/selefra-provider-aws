package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRoute53TrafficPolicyVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRoute53TrafficPolicyVersionsGenerator{}

func (x *TableAwsRoute53TrafficPolicyVersionsGenerator) GetTableName() string {
	return "aws_route53_traffic_policy_versions"
}

func (x *TableAwsRoute53TrafficPolicyVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRoute53TrafficPolicyVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRoute53TrafficPolicyVersionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"traffic_policy_arn",
		},
	}
}

func (x *TableAwsRoute53TrafficPolicyVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.TrafficPolicySummary)
			config := route53.ListTrafficPolicyVersionsInput{Id: r.Id}
			svc := client.(*aws_client.Client).AwsServices().Route53
			for {
				response, err := svc.ListTrafficPolicyVersions(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.TrafficPolicies
				if aws.ToString(response.TrafficPolicyVersionMarker) == "" {
					break
				}
				config.TrafficPolicyVersionMarker = response.TrafficPolicyVersionMarker
			}
			return nil
		},
	}
}

func (x *TableAwsRoute53TrafficPolicyVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsRoute53TrafficPolicyVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("traffic_policy_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Comment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("document").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Document")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_route53_traffic_policies_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_route53_traffic_policies.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsRoute53TrafficPolicyVersionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
