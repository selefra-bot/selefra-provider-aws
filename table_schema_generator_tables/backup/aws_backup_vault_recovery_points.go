package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsBackupVaultRecoveryPointsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsBackupVaultRecoveryPointsGenerator{}

func (x *TableAwsBackupVaultRecoveryPointsGenerator) GetTableName() string {
	return "aws_backup_vault_recovery_points"
}

func (x *TableAwsBackupVaultRecoveryPointsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsBackupVaultRecoveryPointsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsBackupVaultRecoveryPointsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsBackupVaultRecoveryPointsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Backup
			vault := task.ParentRawResult.(types.BackupVaultListMember)
			params := backup.ListRecoveryPointsByBackupVaultInput{BackupVaultName: vault.BackupVaultName, MaxResults: aws.Int32(100)}
			for {
				result, err := svc.ListRecoveryPointsByBackupVault(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.RecoveryPoints
				if aws.ToString(result.NextToken) == "" {
					break
				}
				params.NextToken = result.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsBackupVaultRecoveryPointsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("backup")
}

func (x *TableAwsBackupVaultRecoveryPointsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("composite_member_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CompositeMemberIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreatedBy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_restore_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastRestoreTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_vault_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BackupVaultName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_backup_vaults_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_backup_vaults.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vault_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_encrypted").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsEncrypted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("completion_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CompletionDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recovery_point_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RecoveryPointArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_size_in_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("BackupSizeInBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_vault_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BackupVaultArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Lifecycle")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_backup_vault_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceBackupVaultArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IamRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_parent").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsParent")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RecoveryPointArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_key_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EncryptionKeyArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("calculated_lifecycle").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CalculatedLifecycle")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent_recovery_point_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ParentRecoveryPointArn")).Build(),
	}
}

func (x *TableAwsBackupVaultRecoveryPointsGenerator) GetSubTables() []*schema.Table {
	return nil
}
