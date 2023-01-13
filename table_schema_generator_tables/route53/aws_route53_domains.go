package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRoute53DomainsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRoute53DomainsGenerator{}

func (x *TableAwsRoute53DomainsGenerator) GetTableName() string {
	return "aws_route53_domains"
}

func (x *TableAwsRoute53DomainsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRoute53DomainsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRoute53DomainsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
		},
	}
}

func (x *TableAwsRoute53DomainsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Route53domains
			var input route53domains.ListDomainsInput

			for {
				output, err := svc.ListDomains(ctx, &input, domainClientOpts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.Domains, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Route53domains
					v := result.(types.DomainSummary)

					d, err := svc.GetDomainDetail(ctx, &route53domains.GetDomainDetailInput{DomainName: v.DomainName}, domainClientOpts)
					if err != nil {

					}
					return d, nil

				})
				if aws.ToString(output.NextPageMarker) == "" {
					break
				}
				input.Marker = output.NextPageMarker
			}
			return nil
		},
	}
}

func domainClientOpts(options *route53domains.Options) {

	options.Region = "us-east-1"
}

func (x *TableAwsRoute53DomainsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsRoute53DomainsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("registrar_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegistrarName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("abuse_contact_phone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AbuseContactPhone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expiration_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ExpirationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("who_is_server").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WhoIsServer")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_list").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("StatusList")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transfer_lock").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("admin_privacy").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AdminPrivacy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_sec").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DnsSec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("`A list of tags`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registrant_contact").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RegistrantContact")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registrar_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegistrarUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("abuse_contact_email").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AbuseContactEmail")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nameservers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Nameservers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reseller").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Reseller")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tech_contact").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TechContact")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dnssec_keys").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DnssecKeys")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registrant_privacy").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("RegistrantPrivacy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("UpdatedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResultMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("admin_contact").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AdminContact")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_renew").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AutoRenew")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registry_domain_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegistryDomainId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tech_privacy").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("TechPrivacy")).Build(),
	}
}

func (x *TableAwsRoute53DomainsGenerator) GetSubTables() []*schema.Table {
	return nil
}
