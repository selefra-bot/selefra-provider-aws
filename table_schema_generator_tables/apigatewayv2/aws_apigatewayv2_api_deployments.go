package apigatewayv2

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApigatewayv2ApiDeploymentsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayv2ApiDeploymentsGenerator{}

func (x *TableAwsApigatewayv2ApiDeploymentsGenerator) GetTableName() string {
	return "aws_apigatewayv2_api_deployments"
}

func (x *TableAwsApigatewayv2ApiDeploymentsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayv2ApiDeploymentsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayv2ApiDeploymentsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsApigatewayv2ApiDeploymentsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.Api)
			config := apigatewayv2.GetDeploymentsInput{
				ApiId: r.ApiId,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigatewayv2
			for {
				response, err := svc.GetDeployments(ctx, &config)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Items
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsApigatewayv2ApiDeploymentsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayv2ApiDeploymentsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("api_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				idsComputer := func() ([]string, error) {
					r := result.(types.Deployment)
					p := task.ParentRawResult.(types.Api)
					return []string{"/apis",

						*p.ApiId, "deployments", *r.DeploymentId}, nil
				}

				ids, err := idsComputer()
				if err != nil {
					return nil, diagnostics.AddErrorColumnValueExtractor(task.Table, column, err)
				}

				cl := client.(*aws_client.Client)
				return arn.ARN{
					Partition:	cl.Partition,
					Service:	"apigateway",
					Region:		cl.Region,
					AccountID:	"",
					Resource:	strings.Join(ids, "/"),
				}.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_deployed").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AutoDeployed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeploymentId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeploymentStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment_status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeploymentStatusMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigatewayv2_apis_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigatewayv2_apis.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsApigatewayv2ApiDeploymentsGenerator) GetSubTables() []*schema.Table {
	return nil
}
