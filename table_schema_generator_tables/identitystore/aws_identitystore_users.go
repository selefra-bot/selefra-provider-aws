package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIdentitystoreUsersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIdentitystoreUsersGenerator{}

func (x *TableAwsIdentitystoreUsersGenerator) GetTableName() string {
	return "aws_identitystore_users"
}

func (x *TableAwsIdentitystoreUsersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIdentitystoreUsersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIdentitystoreUsersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsIdentitystoreUsersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			instance, err := getIamInstance(ctx, client)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			svc := client.(*aws_client.Client).AwsServices().Identitystore
			config := identitystore.ListUsersInput{}
			config.IdentityStoreId = instance.IdentityStoreId
			for {
				response, err := svc.ListUsers(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Users

				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsIdentitystoreUsersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("identitystore")
}

func (x *TableAwsIdentitystoreUsersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsIdentitystoreUsersGenerator) GetSubTables() []*schema.Table {
	return nil
}
