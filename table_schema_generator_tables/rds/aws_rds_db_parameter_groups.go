package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRdsDbParameterGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRdsDbParameterGroupsGenerator{}

func (x *TableAwsRdsDbParameterGroupsGenerator) GetTableName() string {
	return "aws_rds_db_parameter_groups"
}

func (x *TableAwsRdsDbParameterGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRdsDbParameterGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRdsDbParameterGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRdsDbParameterGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Rds
			var input rds.DescribeDBParameterGroupsInput
			for {
				output, err := svc.DescribeDBParameterGroups(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.DBParameterGroups
				if aws.ToString(output.Marker) == "" {
					break
				}
				input.Marker = output.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsRdsDbParameterGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("rds")
}

func (x *TableAwsRdsDbParameterGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBParameterGroupArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_parameter_group_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBParameterGroupArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_parameter_group_family").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBParameterGroupFamily")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_parameter_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBParameterGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
	}
}

func (x *TableAwsRdsDbParameterGroupsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsRdsDbParameterGroupDbParametersGenerator{}),
	}
}
