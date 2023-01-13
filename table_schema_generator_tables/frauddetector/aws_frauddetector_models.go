package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsFrauddetectorModelsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsFrauddetectorModelsGenerator{}

func (x *TableAwsFrauddetectorModelsGenerator) GetTableName() string {
	return "aws_frauddetector_models"
}

func (x *TableAwsFrauddetectorModelsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsFrauddetectorModelsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsFrauddetectorModelsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsFrauddetectorModelsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := frauddetector.NewGetModelsPaginator(client.(*aws_client.Client).AwsServices().Frauddetector, nil)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Models
			}
			return nil
		},
	}
}

func (x *TableAwsFrauddetectorModelsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("frauddetector")
}

func (x *TableAwsFrauddetectorModelsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_type_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventTypeName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastUpdatedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ModelId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ModelType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsFrauddetectorModelsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsFrauddetectorModelVersionsGenerator{}),
	}
}
