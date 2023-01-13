package elbv1

import (
	"context"

	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElbv1LoadBalancerPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElbv1LoadBalancerPoliciesGenerator{}

func (x *TableAwsElbv1LoadBalancerPoliciesGenerator) GetTableName() string {
	return "aws_elbv1_load_balancer_policies"
}

func (x *TableAwsElbv1LoadBalancerPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElbv1LoadBalancerPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElbv1LoadBalancerPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsElbv1LoadBalancerPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(ELBv1LoadBalancerWrapper)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Elasticloadbalancing
			response, err := svc.DescribeLoadBalancerPolicies(ctx, &elbv1.DescribeLoadBalancerPoliciesInput{LoadBalancerName: r.LoadBalancerName})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response.PolicyDescriptions
			return nil
		},
	}
}

type ELBv1LoadBalancerWrapper struct {
	types.LoadBalancerDescription
	Tags		map[string]string
	Attributes	*types.LoadBalancerAttributes
}

func (x *TableAwsElbv1LoadBalancerPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticloadbalancing")
}

func (x *TableAwsElbv1LoadBalancerPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_attribute_descriptions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PolicyAttributeDescriptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PolicyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_type_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PolicyTypeName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancer_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancer_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return task.ParentRawResult.(ELBv1LoadBalancerWrapper).LoadBalancerName, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_elbv1_load_balancers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_elbv1_load_balancers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsElbv1LoadBalancerPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
