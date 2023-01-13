package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIamPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamPoliciesGenerator{}

func (x *TableAwsIamPoliciesGenerator) GetTableName() string {
	return "aws_iam_policies"
}

func (x *TableAwsIamPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"id",
		},
	}
}

func (x *TableAwsIamPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			config := iam.GetAccountAuthorizationDetailsInput{
				Filter: []types.EntityType{
					types.EntityTypeAWSManagedPolicy, types.EntityTypeLocalManagedPolicy,
				},
			}
			svc := client.(*aws_client.Client).AwsServices().IAM
			for {
				response, err := svc.GetAccountAuthorizationDetails(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Policies
				if aws.ToString(response.Marker) == "" {
					break
				}
				config.Marker = response.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsIamPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PolicyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("UpdateDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_attachable").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsAttachable")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PolicyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreateDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_version_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultVersionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Path")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions_boundary_usage_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("PermissionsBoundaryUsageCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PolicyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_version_list").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PolicyVersionList")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attachment_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AttachmentCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsIamPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
