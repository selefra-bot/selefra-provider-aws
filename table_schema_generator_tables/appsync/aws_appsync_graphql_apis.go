package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAppsyncGraphqlApisGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAppsyncGraphqlApisGenerator{}

func (x *TableAwsAppsyncGraphqlApisGenerator) GetTableName() string {
	return "aws_appsync_graphql_apis"
}

func (x *TableAwsAppsyncGraphqlApisGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAppsyncGraphqlApisGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAppsyncGraphqlApisGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsAppsyncGraphqlApisGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config appsync.ListGraphqlApisInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Appsync
			for {
				output, err := svc.ListGraphqlApis(ctx, &config, func(options *appsync.Options) {
					options.Region = c.Region
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.GraphqlApis
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsAppsyncGraphqlApisGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("appsync")
}

func (x *TableAwsAppsyncGraphqlApisGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uris").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Uris")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_pool_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UserPoolConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_authentication_providers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AdditionalAuthenticationProviders")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("waf_web_acl_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WafWebAclArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authentication_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthenticationType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LogConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("open_id_connect_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OpenIDConnectConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("xray_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("XrayEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lambda_authorizer_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LambdaAuthorizerConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApiId")).Build(),
	}
}

func (x *TableAwsAppsyncGraphqlApisGenerator) GetSubTables() []*schema.Table {
	return nil
}
