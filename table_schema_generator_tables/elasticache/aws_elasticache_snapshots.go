package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticacheSnapshotsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticacheSnapshotsGenerator{}

func (x *TableAwsElasticacheSnapshotsGenerator) GetTableName() string {
	return "aws_elasticache_snapshots"
}

func (x *TableAwsElasticacheSnapshotsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticacheSnapshotsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticacheSnapshotsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsElasticacheSnapshotsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := elasticache.NewDescribeSnapshotsPaginator(client.(*aws_client.Client).AwsServices().Elasticache, nil)
			for paginator.HasMorePages() {
				v, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- v.Snapshots
			}
			return nil
		},
	}
}

func (x *TableAwsElasticacheSnapshotsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticache")
}

func (x *TableAwsElasticacheSnapshotsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_group_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReplicationGroupDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReplicationGroupId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_maintenance_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredMaintenanceWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_cache_nodes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumCacheNodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshotStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshotWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_minor_version_upgrade").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AutoMinorVersionUpgrade")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("automatic_failover").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AutomaticFailover")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_node_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheNodeType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_subnet_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheSubnetGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Engine")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_availability_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredAvailabilityZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_cluster_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheClusterId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_tiering").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DataTiering")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_snapshots").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NodeSnapshots")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_node_groups").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumNodeGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Port")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_outpost_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredOutpostArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_retention_limit").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SnapshotRetentionLimit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_cluster_create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CacheClusterCreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_parameter_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CacheParameterGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshotName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_source").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshotSource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("topic_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TopicArn")).Build(),
	}
}

func (x *TableAwsElasticacheSnapshotsGenerator) GetSubTables() []*schema.Table {
	return nil
}
