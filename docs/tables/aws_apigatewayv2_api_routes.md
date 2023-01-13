# Table: aws_apigatewayv2_api_routes

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| authorization_scopes | string_array | X | √ |  | 
| route_response_selection_expression | string | X | √ |  | 
| account_id | string | X | √ |  | 
| api_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| route_key | string | X | √ |  | 
| authorizer_id | string | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| api_gateway_managed | bool | X | √ |  | 
| authorization_type | string | X | √ |  | 
| target | string | X | √ |  | 
| request_parameters | json | X | √ |  | 
| route_id | string | X | √ |  | 
| region | string | X | √ |  | 
| api_arn | string | X | √ |  | 
| api_key_required | bool | X | √ |  | 
| model_selection_expression | string | X | √ |  | 
| operation_name | string | X | √ |  | 
| request_models | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


