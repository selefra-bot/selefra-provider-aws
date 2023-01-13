package mq

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsMqBrokerConfigurationRevisionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsMqBrokerConfigurationRevisionsGenerator{}

func (x *TableAwsMqBrokerConfigurationRevisionsGenerator) GetTableName() string {
	return "aws_mq_broker_configuration_revisions"
}

func (x *TableAwsMqBrokerConfigurationRevisionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsMqBrokerConfigurationRevisionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsMqBrokerConfigurationRevisionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsMqBrokerConfigurationRevisionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cfg := task.ParentRawResult.(mq.DescribeConfigurationOutput)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Mq

			input := mq.ListConfigurationRevisionsInput{ConfigurationId: cfg.Id}
			for {
				output, err := svc.ListConfigurationRevisions(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.Revisions, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Mq
					rev := result.(types.ConfigurationRevision)
					cfg := task.ParentRawResult.(mq.DescribeConfigurationOutput)

					revId := strconv.Itoa(int(rev.Revision))
					output, err := svc.DescribeConfigurationRevision(ctx, &mq.DescribeConfigurationRevisionInput{ConfigurationId: cfg.Id, ConfigurationRevision: &revId})
					if err != nil {
						return nil, err
					}
					return output, nil

				})
				if aws.ToString(output.NextToken) == "" {
					break
				}
				input.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsMqBrokerConfigurationRevisionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("mq")
}

func (x *TableAwsMqBrokerConfigurationRevisionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("data").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Data")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResultMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("broker_configuration_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ConfigurationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_mq_broker_configurations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_mq_broker_configurations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsMqBrokerConfigurationRevisionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
