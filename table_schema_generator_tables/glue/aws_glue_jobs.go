package glue

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsGlueJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlueJobsGenerator{}

func (x *TableAwsGlueJobsGenerator) GetTableName() string {
	return "aws_glue_jobs"
}

func (x *TableAwsGlueJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlueJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlueJobsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsGlueJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Glue
			input := glue.GetJobsInput{}
			for {
				result, err := svc.GetJobs(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.Jobs
				if aws.ToString(result.NextToken) == "" {
					break
				}
				input.NextToken = result.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsGlueJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glue")
}

func (x *TableAwsGlueJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedOn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("command").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Command")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("code_gen_configuration_nodes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CodeGenConfigurationNodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("execution_class").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExecutionClass")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_retries").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxRetries")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("non_overridable_arguments").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NonOverridableArguments")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_configuration").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SecurityConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("glue_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GlueVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_control_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SourceControlDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connections").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Connections")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_capacity").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("MaxCapacity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timeout").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Timeout")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notification_property").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NotificationProperty")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("worker_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WorkerType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Role")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allocated_capacity").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AllocatedCapacity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_arguments").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DefaultArguments")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("execution_property").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExecutionProperty")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LogUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number_of_workers").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumberOfWorkers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					return jobARN(cl, aws.ToString(result.(types.Job).Name)), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_on").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastModifiedOn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func jobARN(cl *aws_client.Client, name string) string {
	return arn.ARN{
		Partition:	cl.Partition,
		Service:	"glue",
		Region:		cl.Region,
		AccountID:	cl.AccountID,
		Resource:	fmt.Sprintf("job/%s", name),
	}.String()
}

func (x *TableAwsGlueJobsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsGlueJobRunsGenerator{}),
	}
}
