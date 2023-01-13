package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsGlueDatacatalogEncryptionSettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlueDatacatalogEncryptionSettingsGenerator{}

func (x *TableAwsGlueDatacatalogEncryptionSettingsGenerator) GetTableName() string {
	return "aws_glue_datacatalog_encryption_settings"
}

func (x *TableAwsGlueDatacatalogEncryptionSettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlueDatacatalogEncryptionSettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlueDatacatalogEncryptionSettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
		},
	}
}

func (x *TableAwsGlueDatacatalogEncryptionSettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Glue
			result, err := svc.GetDataCatalogEncryptionSettings(ctx, &glue.GetDataCatalogEncryptionSettingsInput{})
			if err != nil {
				if cl.IsNotFoundError(err) {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- result.DataCatalogEncryptionSettings
			return nil
		},
	}
}

func (x *TableAwsGlueDatacatalogEncryptionSettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glue")
}

func (x *TableAwsGlueDatacatalogEncryptionSettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_password_encryption").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ConnectionPasswordEncryption")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_at_rest").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionAtRest")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsGlueDatacatalogEncryptionSettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
