package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsFsxFileSystemsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsFsxFileSystemsGenerator{}

func (x *TableAwsFsxFileSystemsGenerator) GetTableName() string {
	return "aws_fsx_file_systems"
}

func (x *TableAwsFsxFileSystemsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsFsxFileSystemsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsFsxFileSystemsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsFsxFileSystemsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Fsx
			input := fsx.DescribeFileSystemsInput{MaxResults: aws.Int32(1000)}
			paginator := fsx.NewDescribeFileSystemsPaginator(svc, &input)
			for paginator.HasMorePages() {
				result, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.FileSystems
			}
			return nil
		},
	}
}

func (x *TableAwsFsxFileSystemsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("fsx")
}

func (x *TableAwsFsxFileSystemsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("administrative_actions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AdministrativeActions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StorageType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failure_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FailureDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_system_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FileSystemType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lustre_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LustreConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_interface_ids").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("NetworkInterfaceIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_capacity").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("StorageCapacity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_ids").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SubnetIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("windows_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("WindowsConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_system_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FileSystemId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_system_type_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FileSystemTypeVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OwnerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DNSName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Lifecycle")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ontap_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OntapConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("open_zfs_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OpenZFSConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
	}
}

func (x *TableAwsFsxFileSystemsGenerator) GetSubTables() []*schema.Table {
	return nil
}
