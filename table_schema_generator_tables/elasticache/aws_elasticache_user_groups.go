package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticacheUserGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticacheUserGroupsGenerator{}

func (x *TableAwsElasticacheUserGroupsGenerator) GetTableName() string {
	return "aws_elasticache_user_groups"
}

func (x *TableAwsElasticacheUserGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticacheUserGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticacheUserGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsElasticacheUserGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := elasticache.NewDescribeUserGroupsPaginator(client.(*aws_client.Client).AwsServices().Elasticache, nil)
			for paginator.HasMorePages() {
				v, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- v.UserGroups
			}
			return nil
		},
	}
}

func (x *TableAwsElasticacheUserGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticache")
}

func (x *TableAwsElasticacheUserGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("pending_changes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PendingChanges")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UserGroupId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_ids").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("UserIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_groups").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ReplicationGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Engine")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum_engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MinimumEngineVersion")).Build(),
	}
}

func (x *TableAwsElasticacheUserGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
