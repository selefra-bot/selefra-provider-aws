package apigateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApigatewayRestApisGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayRestApisGenerator{}

func (x *TableAwsApigatewayRestApisGenerator) GetTableName() string {
	return "aws_apigateway_rest_apis"
}

func (x *TableAwsApigatewayRestApisGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayRestApisGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayRestApisGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsApigatewayRestApisGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config apigateway.GetRestApisInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigateway
			for p := apigateway.NewGetRestApisPaginator(svc, &config); p.HasMorePages(); {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Items
			}
			return nil
		},
	}
}

func (x *TableAwsApigatewayRestApisGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayRestApisGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("binary_media_types").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("BinaryMediaTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum_compression_size").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MinimumCompressionSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Policy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disable_execute_api_endpoint").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DisableExecuteApiEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndpointConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("warnings").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Warnings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					rapi := result.(types.RestApi)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	string("apigateway"),
						Region:		cl.Region,
						AccountID:	"",
						Resource:	fmt.Sprintf("/restapis/%s", aws.ToString(rapi.Id)),
					}.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_key_source").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApiKeySource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedDate")).Build(),
	}
}

func (x *TableAwsApigatewayRestApisGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsApigatewayRestApiDocumentationVersionsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayRestApiRequestValidatorsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayRestApiAuthorizersGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayRestApiDeploymentsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayRestApiModelsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayRestApiResourcesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayRestApiStagesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayRestApiDocumentationPartsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayRestApiGatewayResponsesGenerator{}),
	}
}
