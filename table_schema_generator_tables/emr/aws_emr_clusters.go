package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEmrClustersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEmrClustersGenerator{}

func (x *TableAwsEmrClustersGenerator) GetTableName() string {
	return "aws_emr_clusters"
}

func (x *TableAwsEmrClustersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEmrClustersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEmrClustersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEmrClustersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			config := emr.ListClustersInput{
				ClusterStates: []types.ClusterState{
					types.ClusterStateRunning,
					types.ClusterStateStarting,
					types.ClusterStateBootstrapping,
					types.ClusterStateWaiting,
				},
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Emr
			for {
				response, err := svc.ListClusters(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.Clusters, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Emr
					response, err := svc.DescribeCluster(ctx, &emr.DescribeClusterInput{ClusterId: result.(types.ClusterSummary).Id})
					if err != nil {
						return nil, err
					}
					return response.Cluster, nil

				})
				if aws.ToString(response.Marker) == "" {
					break
				}
				config.Marker = response.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsEmrClustersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticmapreduce")
}

func (x *TableAwsEmrClustersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("log_encryption_kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LogEncryptionKmsKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("step_concurrency_level").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("StepConcurrencyLevel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configurations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Configurations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_ami_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustomAmiId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ec2_instance_attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Ec2InstanceAttributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("applications").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Applications")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("outpost_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OutpostArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("placement_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PlacementGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_configuration").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SecurityConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_terminate").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AutoTerminate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ebs_root_volume_size").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("EbsRootVolumeSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("os_release_label").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OSReleaseLabel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("release_label").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReleaseLabel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("running_ami_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RunningAmiVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scale_down_behavior").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ScaleDownBehavior")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_scaling_role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AutoScalingRole")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kerberos_attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("KerberosAttributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceRole")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LogUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_public_dns_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterPublicDnsName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requested_ami_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RequestedAmiVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("termination_protected").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("TerminationProtected")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("visible_to_all_users").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("VisibleToAllUsers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_collection_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceCollectionType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("normalized_instance_hours").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NormalizedInstanceHours")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repo_upgrade_on_boot").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RepoUpgradeOnBoot")).Build(),
	}
}

func (x *TableAwsEmrClustersGenerator) GetSubTables() []*schema.Table {
	return nil
}
