package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIdentitystoreGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIdentitystoreGroupsGenerator{}

func (x *TableAwsIdentitystoreGroupsGenerator) GetTableName() string {
	return "aws_identitystore_groups"
}

func (x *TableAwsIdentitystoreGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIdentitystoreGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIdentitystoreGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsIdentitystoreGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			instance, err := getIamInstance(ctx, client)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			svc := client.(*aws_client.Client).AwsServices().Identitystore
			config := identitystore.ListGroupsInput{}
			config.IdentityStoreId = instance.IdentityStoreId
			for {
				response, err := svc.ListGroups(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Groups

				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func getIamInstance(ctx context.Context, client any) (types.InstanceMetadata, error) {
	svc := client.(*aws_client.Client).AwsServices().Ssoadmin
	config := ssoadmin.ListInstancesInput{}
	response, err := svc.ListInstances(ctx, &config)
	if err == nil {
		for _, i := range response.Instances {
			return i, err
		}
	}
	return types.InstanceMetadata{}, err
}

func (x *TableAwsIdentitystoreGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("identitystore")
}

func (x *TableAwsIdentitystoreGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsIdentitystoreGroupsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsIdentitystoreGroupMembershipsGenerator{}),
	}
}
