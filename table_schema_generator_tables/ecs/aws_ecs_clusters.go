package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEcsClustersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEcsClustersGenerator{}

func (x *TableAwsEcsClustersGenerator) GetTableName() string {
	return "aws_ecs_clusters"
}

func (x *TableAwsEcsClustersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEcsClustersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEcsClustersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEcsClustersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config ecs.ListClustersInput
			svc := client.(*aws_client.Client).AwsServices().Ecs
			for {
				listClustersOutput, err := svc.ListClusters(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				if len(listClustersOutput.ClusterArns) == 0 {
					return nil
				}
				describeClusterOutput, err := svc.DescribeClusters(ctx, &ecs.DescribeClustersInput{
					Clusters:	listClustersOutput.ClusterArns,
					Include:	[]types.ClusterField{types.ClusterFieldTags},
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- describeClusterOutput.Clusters

				if listClustersOutput.NextToken == nil {
					break
				}
				config.NextToken = listClustersOutput.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEcsClustersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ecs")
}

func (x *TableAwsEcsClustersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registered_container_instances_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RegisteredContainerInstancesCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attachments").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Attachments")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Configuration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_capacity_provider_strategy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DefaultCapacityProviderStrategy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active_services_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ActiveServicesCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("running_tasks_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RunningTasksCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_connect_defaults").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServiceConnectDefaults")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attachments_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AttachmentsStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity_providers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("CapacityProviders")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_tasks_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("PendingTasksCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Settings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("statistics").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Statistics")).Build(),
	}
}

func (x *TableAwsEcsClustersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsEcsClusterContainerInstancesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsEcsClusterTasksGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsEcsClusterServicesGenerator{}),
	}
}
