package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCloudfrontDistributionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudfrontDistributionsGenerator{}

func (x *TableAwsCloudfrontDistributionsGenerator) GetTableName() string {
	return "aws_cloudfront_distributions"
}

func (x *TableAwsCloudfrontDistributionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudfrontDistributionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudfrontDistributionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsCloudfrontDistributionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config cloudfront.ListDistributionsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudfront
			for {
				response, err := svc.ListDistributions(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.DistributionList.Items, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Cloudfront

					d := result.(types.DistributionSummary)

					distribution, err := svc.GetDistribution(ctx, &cloudfront.GetDistributionInput{
						Id: d.Id,
					})
					if err != nil {
						return nil, err
					}
					return distribution.Distribution, nil

				})
				if aws.ToString(response.DistributionList.Marker) == "" {
					break
				}
				config.Marker = response.DistributionList.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsCloudfrontDistributionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsCloudfrontDistributionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("distribution_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DistributionConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alias_icp_recordals").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AliasICPRecordals")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("in_progress_invalidation_batches").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("InProgressInvalidationBatches")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastModifiedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active_trusted_key_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ActiveTrustedKeyGroups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active_trusted_signers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ActiveTrustedSigners")).Build(),
	}
}

func (x *TableAwsCloudfrontDistributionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
