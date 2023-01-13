package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticacheGlobalReplicationGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticacheGlobalReplicationGroupsGenerator{}

func (x *TableAwsElasticacheGlobalReplicationGroupsGenerator) GetTableName() string {
	return "aws_elasticache_global_replication_groups"
}

func (x *TableAwsElasticacheGlobalReplicationGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticacheGlobalReplicationGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticacheGlobalReplicationGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsElasticacheGlobalReplicationGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := elasticache.NewDescribeGlobalReplicationGroupsPaginator(client.(*aws_client.Client).AwsServices().Elasticache, nil)
			for paginator.HasMorePages() {
				v, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- v.GlobalReplicationGroups
			}
			return nil
		},
	}
}

func (x *TableAwsElasticacheGlobalReplicationGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticache")
}

func (x *TableAwsElasticacheGlobalReplicationGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("global_replication_group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GlobalReplicationGroupId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Members")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transit_encryption_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("TransitEncryptionEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auth_token_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AuthTokenEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("global_replication_group_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GlobalReplicationGroupDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_node_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheNodeType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Engine")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("at_rest_encryption_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AtRestEncryptionEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ClusterEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("global_node_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("GlobalNodeGroups")).Build(),
	}
}

func (x *TableAwsElasticacheGlobalReplicationGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
