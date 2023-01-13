package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsTransferServersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsTransferServersGenerator{}

func (x *TableAwsTransferServersGenerator) GetTableName() string {
	return "aws_transfer_servers"
}

func (x *TableAwsTransferServersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsTransferServersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsTransferServersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsTransferServersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Transfer
			input := transfer.ListServersInput{MaxResults: aws.Int32(1000)}
			for {
				result, err := svc.ListServers(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, result.Servers, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Transfer
					server := result.(types.ListedServer)

					desc, err := svc.DescribeServer(ctx, &transfer.DescribeServerInput{ServerId: server.ServerId})
					if err != nil {
						return nil, err
					}
					return desc.Server, nil

				})
				if aws.ToString(result.NextToken) == "" {
					break
				}
				input.NextToken = result.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsTransferServersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("transfer")
}

func (x *TableAwsTransferServersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EndpointType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host_key_fingerprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HostKeyFingerprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_provider_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IdentityProviderDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logging_role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LoggingRole")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("protocol_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ProtocolDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("workflow_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("WorkflowDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Certificate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_provider_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IdentityProviderType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pre_authentication_login_banner").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreAuthenticationLoginBanner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("server_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("post_authentication_login_banner").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PostAuthenticationLoginBanner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("UserCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_policy_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SecurityPolicyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Domain")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndpointDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("protocols").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Protocols")).Build(),
	}
}

func (x *TableAwsTransferServersGenerator) GetSubTables() []*schema.Table {
	return nil
}
