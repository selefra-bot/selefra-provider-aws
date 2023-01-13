package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEcrRepositoryImagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEcrRepositoryImagesGenerator{}

func (x *TableAwsEcrRepositoryImagesGenerator) GetTableName() string {
	return "aws_ecr_repository_images"
}

func (x *TableAwsEcrRepositoryImagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEcrRepositoryImagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEcrRepositoryImagesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEcrRepositoryImagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			config := ecr.DescribeImagesInput{
				RepositoryName:	task.ParentRawResult.(types.Repository).RepositoryName,
				MaxResults:	aws.Int32(1000),
			}
			paginator := ecr.NewDescribeImagesPaginator(client.(*aws_client.Client).AwsServices().Ecr, &config)
			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.ImageDetails
			}
			return nil
		},
	}
}

func (x *TableAwsEcrRepositoryImagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsEcrRepositoryImagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("image_scan_findings_summary").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ImageScanFindingsSummary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_scan_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ImageScanStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_size_in_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ImageSizeInBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_ecr_repositories_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_ecr_repositories.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					item := result.(types.ImageDetail)

					a := arn.ARN{
						Partition:	cl.Partition,
						Service:	"ecr",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	"repository_image/" + *item.RegistryId + "/" + *item.ImageDigest,
					}
					return a.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registry_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegistryId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("artifact_media_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ArtifactMediaType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_digest").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageDigest")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_pushed_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ImagePushedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repository_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RepositoryName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_manifest_media_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImageManifestMediaType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_tags").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ImageTags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_recorded_pull_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastRecordedPullTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsEcrRepositoryImagesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsEcrRepositoryImageScanFindingsGenerator{}),
	}
}
