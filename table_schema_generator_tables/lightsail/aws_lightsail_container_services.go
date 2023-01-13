package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLightsailContainerServicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailContainerServicesGenerator{}

func (x *TableAwsLightsailContainerServicesGenerator) GetTableName() string {
	return "aws_lightsail_container_services"
}

func (x *TableAwsLightsailContainerServicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailContainerServicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailContainerServicesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLightsailContainerServicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input lightsail.GetContainerServicesInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			response, err := svc.GetContainerServices(ctx, &input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response.ContainerServices
			return nil
		},
	}
}

func (x *TableAwsLightsailContainerServicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailContainerServicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_deployment").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CurrentDeployment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_disabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsDisabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("power").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Power")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("power_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PowerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateDomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_detail").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StateDetail")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Url")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("container_service_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ContainerServiceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_registry_access").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PrivateRegistryAccess")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_domain_names").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PublicDomainNames")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scale").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Scale")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("principal_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrincipalArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_deployment").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NextDeployment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
	}
}

func (x *TableAwsLightsailContainerServicesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsLightsailContainerServiceDeploymentsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsLightsailContainerServiceImagesGenerator{}),
	}
}
