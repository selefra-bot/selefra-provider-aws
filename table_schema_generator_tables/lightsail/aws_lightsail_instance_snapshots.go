package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLightsailInstanceSnapshotsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailInstanceSnapshotsGenerator{}

func (x *TableAwsLightsailInstanceSnapshotsGenerator) GetTableName() string {
	return "aws_lightsail_instance_snapshots"
}

func (x *TableAwsLightsailInstanceSnapshotsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailInstanceSnapshotsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailInstanceSnapshotsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLightsailInstanceSnapshotsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input lightsail.GetInstanceSnapshotsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			for {
				response, err := svc.GetInstanceSnapshots(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.InstanceSnapshots
				if aws.ToString(response.NextPageToken) == "" {
					break
				}
				input.PageToken = response.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableAwsLightsailInstanceSnapshotsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailInstanceSnapshotsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("from_instance_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FromInstanceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_in_gb").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SizeInGb")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("from_blueprint_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FromBlueprintId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("from_bundle_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FromBundleId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_from_auto_snapshot").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsFromAutoSnapshot")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("progress").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Progress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("from_instance_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FromInstanceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("from_attached_disks").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FromAttachedDisks")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("support_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SupportCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
	}
}

func (x *TableAwsLightsailInstanceSnapshotsGenerator) GetSubTables() []*schema.Table {
	return nil
}
