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

type TableAwsApigatewayv2ApiIntegrationResponsesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayv2ApiIntegrationResponsesGenerator{}

func (x *TableAwsApigatewayv2ApiIntegrationResponsesGenerator) GetTableName() string {
	return "aws_apigatewayv2_api_integration_responses"
}

func (x *TableAwsApigatewayv2ApiIntegrationResponsesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayv2ApiIntegrationResponsesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayv2ApiIntegrationResponsesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsApigatewayv2ApiIntegrationResponsesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.Integration)
			p := task.ParentTask.ParentRawResult.(types.Api)
			config := apigatewayv2.GetIntegrationResponsesInput{
				ApiId:		p.ApiId,
				IntegrationId:	r.IntegrationId,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigatewayv2
			for {
				response, err := svc.GetIntegrationResponses(ctx, &config)

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

func (x *TableAwsApigatewayv2ApiIntegrationResponsesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayv2ApiIntegrationResponsesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("integration_response_key").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IntegrationResponseKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_handling_strategy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ContentHandlingStrategy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integration_response_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IntegrationResponseId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("response_templates").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResponseTemplates")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigatewayv2_api_integrations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigatewayv2_api_integrations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_integration_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				idsComputer := func() ([]string, error) {
					r := result.(types.IntegrationResponse)
					i := task.ParentRawResult.(types.Integration)
					api := task.ParentTask.ParentRawResult.(types.Api)
					return []string{"/apis",

						*api.ApiId, "integrations",

						*i.IntegrationId, "integrationresponses", *r.IntegrationResponseId}, nil
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
		table_schema_generator.NewColumnBuilder().ColumnName("template_selection_expression").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TemplateSelectionExpression")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integration_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("integration_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("response_parameters").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResponseParameters")).Build(),
	}
}

func (x *TableAwsApigatewayv2ApiIntegrationResponsesGenerator) GetSubTables() []*schema.Table {
	return nil
}
