package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsQldbLedgersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsQldbLedgersGenerator{}

func (x *TableAwsQldbLedgersGenerator) GetTableName() string {
	return "aws_qldb_ledgers"
}

func (x *TableAwsQldbLedgersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsQldbLedgersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsQldbLedgersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsQldbLedgersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Qldb
			config := qldb.ListLedgersInput{}
			for {
				response, err := svc.ListLedgers(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.Ledgers, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Qldb
					l := result.(types.LedgerSummary)

					response, err := svc.DescribeLedger(ctx, &qldb.DescribeLedgerInput{Name: l.Name})
					if err != nil {
						return nil, err
					}
					return response, nil

				})
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsQldbLedgersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("qldb")
}

func (x *TableAwsQldbLedgersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("`The tags associated with the pipeline.`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationDateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_protection").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DeletionProtection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_description").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PermissionsMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResultMetadata")).Build(),
	}
}

func (x *TableAwsQldbLedgersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsQldbLedgerJournalS3ExportsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsQldbLedgerJournalKinesisStreamsGenerator{}),
	}
}
