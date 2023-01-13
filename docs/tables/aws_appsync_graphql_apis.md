# Table: aws_appsync_graphql_apis

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| uris | json | X | √ |  | 
| user_pool_config | json | X | √ |  | 
| region | string | X | √ |  | 
| additional_authentication_providers | json | X | √ |  | 
| tags | json | X | √ |  | 
| waf_web_acl_arn | string | X | √ |  | 
| authentication_type | string | X | √ |  | 
| log_config | json | X | √ |  | 
| open_id_connect_config | json | X | √ |  | 
| xray_enabled | bool | X | √ |  | 
| arn | string | X | √ |  | 
| lambda_authorizer_config | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| api_id | string | X | √ |  | 


