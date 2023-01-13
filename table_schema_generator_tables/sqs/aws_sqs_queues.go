package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/mitchellh/mapstructure"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSqsQueuesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSqsQueuesGenerator{}

func (x *TableAwsSqsQueuesGenerator) GetTableName() string {
	return "aws_sqs_queues"
}

func (x *TableAwsSqsQueuesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSqsQueuesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSqsQueuesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsSqsQueuesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Sqs
			var params sqs.ListQueuesInput
			for {
				result, err := svc.ListQueues(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, result.QueueUrls, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Sqs
					qURL := result.(string)

					input := sqs.GetQueueAttributesInput{
						QueueUrl:	aws.String(qURL),
						AttributeNames:	[]types.QueueAttributeName{types.QueueAttributeNameAll},
					}
					out, err := svc.GetQueueAttributes(ctx, &input)
					if err != nil {
						return nil, err
					}

					q := &Queue{URL: qURL}
					d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: q})
					if err != nil {
						return nil, err
					}
					if err := d.Decode(out.Attributes); err != nil {
						return nil, err
					}
					return q, nil

				})
				if aws.ToString(result.NextToken) == "" {
					break
				}
				params.NextToken = result.NextToken
			}
			return nil
		},
	}
}

type Queue struct {
	URL	string

	ApproximateNumberOfMessages	*int32

	ApproximateNumberOfMessagesDelayed	*int32

	ApproximateNumberOfMessagesNotVisible	*int32

	CreatedTimestamp	*int32

	DelaySeconds	*int32

	LastModifiedTimestamp	*int32

	MaximumMessageSize	*int32

	MessageRetentionPeriod	*int32

	Policy	*string

	Arn	*string	`mapstructure:"QueueArn"`

	ReceiveMessageWaitTimeSeconds	*int32

	RedrivePolicy	*string

	VisibilityTimeout	*int32

	KmsMasterKeyId	*string

	KmsDataKeyReusePeriodSeconds	*int32

	SqsManagedSseEnabled	*bool

	FifoQueue	*bool

	ContentBasedDeduplication	*bool

	DeduplicationScope	*string

	FifoThroughputLimit	*string

	RedriveAllowPolicy	*string

	UnknownFields	map[string]any	`mapstructure:",remain"`
}

func (x *TableAwsSqsQueuesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("sqs")
}

func (x *TableAwsSqsQueuesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Policy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("approximate_number_of_messages").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ApproximateNumberOfMessages")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sqs_managed_sse_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SqsManagedSseEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unknown_fields").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UnknownFields")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_timestamp").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("CreatedTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_timestamp").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("LastModifiedTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_data_key_reuse_period_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("KmsDataKeyReusePeriodSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("approximate_number_of_messages_delayed").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ApproximateNumberOfMessagesDelayed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum_message_size").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaximumMessageSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_master_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsMasterKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_based_deduplication").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ContentBasedDeduplication")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deduplication_scope").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeduplicationScope")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("redrive_allow_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RedriveAllowPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("visibility_timeout").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("VisibilityTimeout")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fifo_queue").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("FifoQueue")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("redrive_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RedrivePolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("approximate_number_of_messages_not_visible").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ApproximateNumberOfMessagesNotVisible")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delay_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DelaySeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fifo_throughput_limit").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FifoThroughputLimit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message_retention_period").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MessageRetentionPeriod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("receive_message_wait_time_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ReceiveMessageWaitTimeSeconds")).Build(),
	}
}

func (x *TableAwsSqsQueuesGenerator) GetSubTables() []*schema.Table {
	return nil
}
