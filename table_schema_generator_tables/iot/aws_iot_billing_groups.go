package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIotBillingGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIotBillingGroupsGenerator{}

func (x *TableAwsIotBillingGroupsGenerator) GetTableName() string {
	return "aws_iot_billing_groups"
}

func (x *TableAwsIotBillingGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIotBillingGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIotBillingGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsIotBillingGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			input := iot.ListBillingGroupsInput{
				MaxResults: aws.Int32(250),
			}
			c := client.(*aws_client.Client)

			svc := c.AwsServices().Iot
			for {
				response, err := svc.ListBillingGroups(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, g := range response.BillingGroups {
					group, err := svc.DescribeBillingGroup(ctx, &iot.DescribeBillingGroupInput{
						BillingGroupName: g.GroupName,
					}, func(options *iot.Options) {
						options.Region = c.Region
					})
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- group
				}

				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsIotBillingGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("iot")
}

func (x *TableAwsIotBillingGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BillingGroupArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing_group_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BillingGroupArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing_group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BillingGroupId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing_group_properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BillingGroupProperties")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResultMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("things_in_group").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					i := result.(*iot.DescribeBillingGroupOutput)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Iot
					input := iot.ListThingsInBillingGroupInput{
						BillingGroupName:	i.BillingGroupName,
						MaxResults:		aws.Int32(250),
					}

					var things []string
					for {
						response, err := svc.ListThingsInBillingGroup(ctx, &input)
						if err != nil {
							return nil, err
						}

						things = append(things, response.Things...)

						if aws.ToString(response.NextToken) == "" {
							break
						}
						input.NextToken = response.NextToken
					}
					return things, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing_group_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BillingGroupMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BillingGroupName")).Build(),
	}
}

func (x *TableAwsIotBillingGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
