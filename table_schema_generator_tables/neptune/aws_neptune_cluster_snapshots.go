package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsNeptuneClusterSnapshotsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsNeptuneClusterSnapshotsGenerator{}

func (x *TableAwsNeptuneClusterSnapshotsGenerator) GetTableName() string {
	return "aws_neptune_cluster_snapshots"
}

func (x *TableAwsNeptuneClusterSnapshotsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsNeptuneClusterSnapshotsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsNeptuneClusterSnapshotsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsNeptuneClusterSnapshotsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Neptune
			input := neptune.DescribeDBClusterSnapshotsInput{
				Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
			}
			for {
				output, err := svc.DescribeDBClusterSnapshots(ctx, &input)
				if err != nil {
					return nil
				}
				resultChannel <- output.DBClusterSnapshots
				if aws.ToString(output.Marker) == "" {
					break
				}
				input.Marker = output.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsNeptuneClusterSnapshotsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("neptune")
}

func (x *TableAwsNeptuneClusterSnapshotsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allocated_storage").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AllocatedStorage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Engine")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("percent_progress").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("PercentProgress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Port")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ClusterCreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_username").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterUsername")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_db_cluster_snapshot_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceDBClusterSnapshotArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterSnapshotArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_snapshot_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterSnapshotArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_snapshot_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterSnapshotIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnapshotType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					s := result.(types.DBClusterSnapshot)
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Neptune
					out, err := svc.DescribeDBClusterSnapshotAttributes(
						ctx,
						&neptune.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
						func(o *neptune.Options) {
							o.Region = c.Region
						},
					)
					if err != nil {
						if c.IsNotFoundError(err) {
							return nil, nil
						}
						return nil, err
					}
					if out.DBClusterSnapshotAttributesResult == nil {
						return nil, nil
					}

					return out.DBClusterSnapshotAttributesResult.DBClusterSnapshotAttributes, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zones").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("AvailabilityZones")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_database_authentication_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IAMDatabaseAuthenticationEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("license_model").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LicenseModel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("SnapshotCreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_encrypted").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("StorageEncrypted")).Build(),
	}
}

func (x *TableAwsNeptuneClusterSnapshotsGenerator) GetSubTables() []*schema.Table {
	return nil
}
