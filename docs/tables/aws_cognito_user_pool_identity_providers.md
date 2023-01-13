# Table: aws_cognito_user_pool_identity_providers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| attribute_mapping | json | X | √ |  | 
| provider_details | json | X | √ |  | 
| provider_name | string | X | √ |  | 
| user_pool_id | string | X | √ |  | 
| aws_cognito_user_pools_selefra_id | string | X | X | fk to aws_cognito_user_pools.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| user_pool_arn | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| idp_identifiers | string_array | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| provider_type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


