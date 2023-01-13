# Table: aws_apigatewayv2_api_authorizers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| api_arn | string | X | √ |  | 
| api_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| authorizer_payload_format_version | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| name | string | X | √ |  | 
| authorizer_id | string | X | √ |  | 
| authorizer_result_ttl_in_seconds | big_int | X | √ |  | 
| enable_simple_responses | bool | X | √ |  | 
| identity_source | string_array | X | √ |  | 
| authorizer_credentials_arn | string | X | √ |  | 
| identity_validation_expression | string | X | √ |  | 
| jwt_configuration | json | X | √ |  | 
| authorizer_type | string | X | √ |  | 
| authorizer_uri | string | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 


