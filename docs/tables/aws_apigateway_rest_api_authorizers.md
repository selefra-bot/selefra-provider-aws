# Table: aws_apigateway_rest_api_authorizers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ |  | 
| authorizer_uri | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| auth_type | string | X | √ |  | 
| authorizer_credentials | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| identity_source | string | X | √ |  | 
| arn | string | X | √ |  | 
| authorizer_result_ttl_in_seconds | big_int | X | √ |  | 
| id | string | X | √ |  | 
| identity_validation_expression | string | X | √ |  | 
| name | string | X | √ |  | 
| provider_ar_ns | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 


