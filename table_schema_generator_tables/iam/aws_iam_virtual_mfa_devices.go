package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIamVirtualMfaDevicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamVirtualMfaDevicesGenerator{}

func (x *TableAwsIamVirtualMfaDevicesGenerator) GetTableName() string {
	return "aws_iam_virtual_mfa_devices"
}

func (x *TableAwsIamVirtualMfaDevicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamVirtualMfaDevicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamVirtualMfaDevicesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsIamVirtualMfaDevicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config iam.ListVirtualMFADevicesInput
			svc := client.(*aws_client.Client).AwsServices().IAM
			for {
				response, err := svc.ListVirtualMFADevices(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.VirtualMFADevices
				if aws.ToString(response.Marker) == "" {
					break
				}
				config.Marker = response.Marker
			}

			return nil
		},
	}
}

func (x *TableAwsIamVirtualMfaDevicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamVirtualMfaDevicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("serial_number").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SerialNumber")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("base32_string_seed").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("Base32StringSeed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("EnableDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("qr_code_png").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("QRCodePNG")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("User")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsIamVirtualMfaDevicesGenerator) GetSubTables() []*schema.Table {
	return nil
}
