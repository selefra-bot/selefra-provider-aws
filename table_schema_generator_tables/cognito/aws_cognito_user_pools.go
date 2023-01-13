package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCognitoUserPoolsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCognitoUserPoolsGenerator{}

func (x *TableAwsCognitoUserPoolsGenerator) GetTableName() string {
	return "aws_cognito_user_pools"
}

func (x *TableAwsCognitoUserPoolsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCognitoUserPoolsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCognitoUserPoolsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
		},
	}
}

func (x *TableAwsCognitoUserPoolsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cognitoidentityprovider
			params := cognitoidentityprovider.ListUserPoolsInput{

				MaxResults: 60,
			}
			for {
				out, err := svc.ListUserPools(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, out.UserPools, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Cognitoidentityprovider
					item := result.(types.UserPoolDescriptionType)

					upo, err := svc.DescribeUserPool(ctx, &cognitoidentityprovider.DescribeUserPoolInput{UserPoolId: item.Id})
					if err != nil {
						return nil, err
					}
					return upo.UserPool, nil

				})
				if aws.ToString(out.NextToken) == "" {
					break
				}
				params.NextToken = out.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsCognitoUserPoolsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("cognito-identity")
}

func (x *TableAwsCognitoUserPoolsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("mfa_configuration").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MfaConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policies").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Policies")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("username_attributes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("UsernameAttributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("verification_message_template").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VerificationMessageTemplate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_protection").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeletionProtection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lambda_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LambdaConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sms_verification_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SmsVerificationMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_domain").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustomDomain")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("device_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DeviceConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_recovery_setting").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AccountRecoverySetting")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sms_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SmsConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sms_configuration_failure").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SmsConfigurationFailure")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_pool_add_ons").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UserPoolAddOns")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_verification_subject").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EmailVerificationSubject")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("admin_create_user_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AdminCreateUserConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_pool_tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UserPoolTags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Domain")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_configuration_failure").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EmailConfigurationFailure")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("estimated_number_of_users").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("EstimatedNumberOfUsers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("LastModifiedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SchemaAttributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("username_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UsernameConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_verified_attributes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("AutoVerifiedAttributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alias_attributes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("AliasAttributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EmailConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_attribute_update_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UserAttributeUpdateSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_verification_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EmailVerificationMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sms_authentication_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SmsAuthenticationMessage")).Build(),
	}
}

func (x *TableAwsCognitoUserPoolsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsCognitoUserPoolIdentityProvidersGenerator{}),
	}
}
