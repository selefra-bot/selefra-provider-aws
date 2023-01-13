package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsWafv2IpsetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsWafv2IpsetsGenerator{}

func (x *TableAwsWafv2IpsetsGenerator) GetTableName() string {
	return "aws_wafv2_ipsets"
}

func (x *TableAwsWafv2IpsetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsWafv2IpsetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsWafv2IpsetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsWafv2IpsetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Wafv2

			params := wafv2.ListIPSetsInput{
				Scope:	cl.WAFScope,
				Limit:	aws.Int32(100),
			}
			for {
				result, err := svc.ListIPSets(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, result.IPSets, func(result any) (any, error) {
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Wafv2
					s := result.(types.IPSetSummary)

					info, err := svc.GetIPSet(
						ctx,
						&wafv2.GetIPSetInput{
							Id:	s.Id,
							Name:	s.Name,
							Scope:	cl.WAFScope,
						},
						func(options *wafv2.Options) {
							options.Region = cl.Region
						},
					)
					if err != nil {
						return nil, err
					}
					return info.IPSet, nil

				})
				if aws.ToString(result.NextMarker) == "" {
					break
				}
				params.NextMarker = result.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsWafv2IpsetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegionAndScope("waf-regional")
}

func (x *TableAwsWafv2IpsetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("addresses").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Addresses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IPAddressVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsWafv2IpsetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
