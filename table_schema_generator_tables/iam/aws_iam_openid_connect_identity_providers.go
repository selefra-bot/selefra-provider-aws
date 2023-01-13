package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIamOpenidConnectIdentityProvidersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamOpenidConnectIdentityProvidersGenerator{}

func (x *TableAwsIamOpenidConnectIdentityProvidersGenerator) GetTableName() string {
	return "aws_iam_openid_connect_identity_providers"
}

func (x *TableAwsIamOpenidConnectIdentityProvidersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamOpenidConnectIdentityProvidersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamOpenidConnectIdentityProvidersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsIamOpenidConnectIdentityProvidersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().IAM
			response, err := svc.ListOpenIDConnectProviders(ctx, &iam.ListOpenIDConnectProvidersInput{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			aws_client.SendResults(resultChannel, response.OpenIDConnectProviderList, func(result any) (any, error) {
				svc := client.(*aws_client.Client).AwsServices().IAM

				p := result.(types.OpenIDConnectProviderListEntry)
				providerResponse, err := svc.GetOpenIDConnectProvider(ctx, &iam.GetOpenIDConnectProviderInput{OpenIDConnectProviderArn: p.Arn})
				if err != nil {
					return nil, err
				}
				return &IamOpenIdIdentityProviderWrapper{GetOpenIDConnectProviderOutput: providerResponse, Arn: *p.Arn}, nil

			})
			return nil
		},
	}
}

type IamOpenIdIdentityProviderWrapper struct {
	*iam.GetOpenIDConnectProviderOutput
	Arn	string
}

func (x *TableAwsIamOpenidConnectIdentityProvidersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamOpenidConnectIdentityProvidersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("get_open_id_connect_provider_output").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("GetOpenIDConnectProviderOutput")).Build(),
	}
}

func (x *TableAwsIamOpenidConnectIdentityProvidersGenerator) GetSubTables() []*schema.Table {
	return nil
}
