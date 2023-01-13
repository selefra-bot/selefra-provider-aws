# Table: aws_cognito_user_pools

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| mfa_configuration | string | X | √ |  | 
| policies | json | X | √ |  | 
| username_attributes | string_array | X | √ |  | 
| verification_message_template | json | X | √ |  | 
| account_id | string | X | √ |  | 
| deletion_protection | string | X | √ |  | 
| lambda_config | json | X | √ |  | 
| sms_verification_message | string | X | √ |  | 
| region | string | X | √ |  | 
| custom_domain | string | X | √ |  | 
| device_configuration | json | X | √ |  | 
| name | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | X | √ |  | 
| account_recovery_setting | json | X | √ |  | 
| sms_configuration | json | X | √ |  | 
| sms_configuration_failure | string | X | √ |  | 
| user_pool_add_ons | json | X | √ |  | 
| arn | string | X | √ |  | 
| email_verification_subject | string | X | √ |  | 
| admin_create_user_config | json | X | √ |  | 
| user_pool_tags | json | X | √ |  | 
| domain | string | X | √ |  | 
| email_configuration_failure | string | X | √ |  | 
| estimated_number_of_users | big_int | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| schema_attributes | json | X | √ |  | 
| username_configuration | json | X | √ |  | 
| auto_verified_attributes | string_array | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| alias_attributes | string_array | X | √ |  | 
| email_configuration | json | X | √ |  | 
| user_attribute_update_settings | json | X | √ |  | 
| email_verification_message | string | X | √ |  | 
| sms_authentication_message | string | X | √ |  | 


