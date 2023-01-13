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

type TableAwsGlueCrawlersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlueCrawlersGenerator{}

func (x *TableAwsGlueCrawlersGenerator) GetTableName() string {
	return "aws_glue_crawlers"
}

func (x *TableAwsGlueCrawlersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlueCrawlersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlueCrawlersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsGlueCrawlersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Glue
			input := glue.GetCrawlersInput{}
			for {
				output, err := svc.GetCrawlers(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Crawlers

				if aws.ToString(output.NextToken) == "" {
					break
				}
				input.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsGlueCrawlersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glue")
}

func (x *TableAwsGlueCrawlersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("crawler_security_configuration").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CrawlerSecurityConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_crawl").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LastCrawl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("targets").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Targets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crawl_elapsed_time").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("CrawlElapsedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastUpdated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lineage_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LineageConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Role")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Schedule")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Configuration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lake_formation_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LakeFormationConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recrawl_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RecrawlPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("classifiers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Classifiers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					return crawlerARN(cl, aws.ToString(result.(types.Crawler).Name)), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DatabaseName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_change_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SchemaChangePolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("table_prefix").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TablePrefix")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func crawlerARN(cl *aws_client.Client, name string) string {
	return arn.ARN{
		Partition:	cl.Partition,
		Service:	"glue",
		Region:		cl.Region,
		AccountID:	cl.AccountID,
		Resource:	fmt.Sprintf("crawler/%s", name),
	}.String()
}

func (x *TableAwsGlueCrawlersGenerator) GetSubTables() []*schema.Table {
	return nil
}
