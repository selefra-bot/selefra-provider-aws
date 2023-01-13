package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLightsailCertificatesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailCertificatesGenerator{}

func (x *TableAwsLightsailCertificatesGenerator) GetTableName() string {
	return "aws_lightsail_certificates"
}

func (x *TableAwsLightsailCertificatesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailCertificatesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailCertificatesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLightsailCertificatesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			input := lightsail.GetCertificatesInput{
				IncludeCertificateDetails: true,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			response, err := svc.GetCertificates(ctx, &input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			for _, cer := range response.Certificates {
				resultChannel <- cer.CertificateDetail
			}
			return nil
		},
	}
}

func (x *TableAwsLightsailCertificatesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailCertificatesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("subject_alternative_names").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SubjectAlternativeNames")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("in_use_resource_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("InUseResourceCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_algorithm").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KeyAlgorithm")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("revocation_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RevocationReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serial_number").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SerialNumber")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issued_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("IssuedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issuer_ca").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IssuerCA")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("request_failure_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RequestFailureReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_validation_records").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DomainValidationRecords")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("not_after").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("NotAfter")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("renewal_summary").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RenewalSummary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("support_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SupportCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("eligible_to_renew").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EligibleToRenew")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("not_before").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("NotBefore")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("revoked_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("RevokedAt")).Build(),
	}
}

func (x *TableAwsLightsailCertificatesGenerator) GetSubTables() []*schema.Table {
	return nil
}
