package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEksClustersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEksClustersGenerator{}

func (x *TableAwsEksClustersGenerator) GetTableName() string {
	return "aws_eks_clusters"
}

func (x *TableAwsEksClustersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEksClustersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEksClustersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsEksClustersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config eks.ListClustersInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Eks
			for {
				listClustersOutput, err := svc.ListClusters(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, listClustersOutput.Clusters, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Eks
					name := result.(string)
					output, err := svc.DescribeCluster(
						ctx, &eks.DescribeClusterInput{Name: &name}, func(options *eks.Options) {
							options.Region = c.Region
						})
					if err != nil {
						return nil, err
					}
					return output.Cluster, nil

				})
				if listClustersOutput.NextToken == nil {
					break
				}
				config.NextToken = listClustersOutput.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEksClustersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("eks")
}

func (x *TableAwsEksClustersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("connector_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ConnectorConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PlatformVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_authority").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CertificateAuthority")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Identity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logging").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Logging")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("outpost_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OutpostConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_request_token").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClientRequestToken")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kubernetes_network_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("KubernetesNetworkConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources_vpc_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourcesVpcConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Health")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Endpoint")).Build(),
	}
}

func (x *TableAwsEksClustersGenerator) GetSubTables() []*schema.Table {
	return nil
}
