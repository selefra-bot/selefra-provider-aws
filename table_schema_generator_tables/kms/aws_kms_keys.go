package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsKmsKeysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsKmsKeysGenerator{}

func (x *TableAwsKmsKeysGenerator) GetTableName() string {
	return "aws_kms_keys"
}

func (x *TableAwsKmsKeysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsKmsKeysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsKmsKeysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsKmsKeysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Kms

			config := kms.ListKeysInput{Limit: aws.Int32(1000)}
			p := kms.NewListKeysPaginator(svc, &config)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.Keys, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Kms
					item := result.(types.KeyListEntry)

					d, err := svc.DescribeKey(ctx, &kms.DescribeKeyInput{KeyId: item.KeyId})
					if err != nil {
						return nil, err
					}
					return d.KeyMetadata, nil

				})
			}
			return nil
		},
	}
}

func (x *TableAwsKmsKeysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("kms")
}

func (x *TableAwsKmsKeysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("customer_master_key_spec").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustomerMasterKeySpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_algorithms").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("EncryptionAlgorithms")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("multi_region").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("MultiRegion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_manager").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KeyManager")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("signing_algorithms").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SigningAlgorithms")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AWSAccountId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("DeletionDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("multi_region_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MultiRegionConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rotation_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Kms
					key := result.(*types.KeyMetadata)
					if key.Origin == "EXTERNAL" || key.KeyManager == "AWS" {
						return nil, nil
					}
					result, err := svc.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{KeyId: key.KeyId})
					if err != nil {
						return nil, err
					}
					return result.KeyRotationEnabled, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Enabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replica_keys").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					key := result.(*types.KeyMetadata)
					if key.MultiRegionConfiguration == nil {
						return nil, nil
					}
					return key.MultiRegionConfiguration.ReplicaKeys, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloud_hsm_cluster_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CloudHsmClusterId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_spec").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KeySpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mac_algorithms").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("MacAlgorithms")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("xks_key_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("XksKeyConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expiration_model").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExpirationModel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KeyState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_deletion_window_in_days").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("PendingDeletionWindowInDays")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_to").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ValidTo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_key_store_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustomKeyStoreId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_usage").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KeyUsage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("origin").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Origin")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsKmsKeysGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsKmsKeyGrantsGenerator{}),
	}
}
