package ecs

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEcsTaskDefinitionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEcsTaskDefinitionsGenerator{}

func (x *TableAwsEcsTaskDefinitionsGenerator) GetTableName() string {
	return "aws_ecs_task_definitions"
}

func (x *TableAwsEcsTaskDefinitionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEcsTaskDefinitionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEcsTaskDefinitionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEcsTaskDefinitionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config ecs.ListTaskDefinitionsInput
			svc := client.(*aws_client.Client).AwsServices().Ecs
			for {
				listClustersOutput, err := svc.ListTaskDefinitions(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, listClustersOutput.TaskDefinitionArns, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Ecs
					taskArn := result.(string)

					describeTaskDefinitionOutput, err := svc.DescribeTaskDefinition(ctx, &ecs.DescribeTaskDefinitionInput{
						TaskDefinition:	aws.String(taskArn),
						Include:	[]types.TaskDefinitionField{types.TaskDefinitionFieldTags},
					})
					if err != nil {
						return nil, err
					}
					if describeTaskDefinitionOutput.TaskDefinition == nil {
						return nil, errors.New("nil TaskDefinition encountered")
					}
					return TaskDefinitionWrapper{
						TaskDefinition:	describeTaskDefinitionOutput.TaskDefinition,
						Tags:		describeTaskDefinitionOutput.Tags,
					}, nil

				})
				if listClustersOutput.NextToken == nil {
					break
				}
				config.NextToken = listClustersOutput.NextToken
			}
			return nil
		},
	}
}

type TaskDefinitionWrapper struct {
	*types.TaskDefinition
	Tags	[]types.Tag
}

func (x *TableAwsEcsTaskDefinitionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ecs")
}

func (x *TableAwsEcsTaskDefinitionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ephemeral_storage").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EphemeralStorage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("execution_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExecutionRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pid_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PidMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inference_accelerators").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("InferenceAccelerators")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("placement_constraints").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PlacementConstraints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registered_by").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegisteredBy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipc_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IpcMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("proxy_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ProxyConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("revision").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Revision")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TaskDefinitionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compatibilities").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Compatibilities")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("family").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Family")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registered_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("RegisteredAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("volumes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Volumes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deregistered_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("DeregisteredAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("task_definition_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TaskDefinitionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("container_definitions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ContainerDefinitions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cpu").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Cpu")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requires_compatibilities").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("RequiresCompatibilities")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("runtime_platform").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RuntimePlatform")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("memory").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Memory")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NetworkMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requires_attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RequiresAttributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("task_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TaskRoleArn")).Build(),
	}
}

func (x *TableAwsEcsTaskDefinitionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
