package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLightsailInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailInstancesGenerator{}

func (x *TableAwsLightsailInstancesGenerator) GetTableName() string {
	return "aws_lightsail_instances"
}

func (x *TableAwsLightsailInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLightsailInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			input := lightsail.GetInstancesInput{}
			for {
				output, err := svc.GetInstances(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Instances

				if aws.ToString(output.NextPageToken) == "" {
					break
				}
				input.PageToken = output.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableAwsLightsailInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("blueprint_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BlueprintName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IpAddressType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_static_ip").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsStaticIp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("support_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SupportCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("blueprint_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BlueprintId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bundle_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BundleId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_addresses").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Ipv6Addresses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("username").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Username")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					r := result.(types.Instance)
					cli := client.(*aws_client.Client)
					svc := cli.AwsServices().Lightsail
					input := lightsail.GetInstanceAccessDetailsInput{InstanceName: r.Name}
					output, err := svc.GetInstanceAccessDetails(ctx, &input)
					if err != nil {
						return nil, err
					}
					return output.AccessDetails, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("networking").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Networking")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateIpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssh_key_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SshKeyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("add_ons").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AddOns")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hardware").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Hardware")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MetadataOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PublicIpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsLightsailInstancesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsLightsailInstancePortStatesGenerator{}),
	}
}
