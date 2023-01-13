package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDocdbClustersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDocdbClustersGenerator{}

func (x *TableAwsDocdbClustersGenerator) GetTableName() string {
	return "aws_docdb_clusters"
}

func (x *TableAwsDocdbClustersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDocdbClustersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDocdbClustersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsDocdbClustersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Docdb

			input := docdb.DescribeDBClustersInput{
				Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"docdb"}}},
			}

			p := docdb.NewDescribeDBClustersPaginator(svc, &input)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.DBClusters
			}
			return nil
		},
	}
}

func (x *TableAwsDocdbClustersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("docdb")
}

func (x *TableAwsDocdbClustersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("multi_az").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("MultiAZ")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reader_endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReaderEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_retention_period").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("BackupRetentionPeriod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Engine")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zones").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("AvailabilityZones")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_source_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReplicationSourceIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hosted_zone_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HostedZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_maintenance_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredMaintenanceWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_encrypted").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("StorageEncrypted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled_cloudwatch_logs_exports").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("EnabledCloudwatchLogsExports")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Port")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_backup_window").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PreferredBackupWindow")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("read_replica_identifiers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ReadReplicaIdentifiers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_username").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterUsername")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("percent_progress").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PercentProgress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_security_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VpcSecurityGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("clone_group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CloneGroupId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_members").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DBClusterMembers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_protection").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DeletionProtection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Endpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("associated_roles").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AssociatedRoles")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_subnet_group").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBSubnetGroup")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_restorable_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LatestRestorableTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ClusterCreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_parameter_group").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterParameterGroup")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DbClusterResourceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("earliest_restorable_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("EarliestRestorableTime")).Build(),
	}
}

func (x *TableAwsDocdbClustersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsDocdbClusterSnapshotsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsDocdbInstancesGenerator{}),
	}
}
