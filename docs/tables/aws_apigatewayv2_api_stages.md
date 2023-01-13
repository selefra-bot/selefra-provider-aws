# Table: aws_apigatewayv2_api_stages

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| access_log_settings | json | X | √ |  | 
| stage_variables | json | X | √ |  | 
| tags | json | X | √ |  | 
| api_id | string | X | √ |  | 
| api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| stage_name | string | X | √ |  | 
| api_gateway_managed | bool | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| default_route_settings | json | X | √ |  | 
| description | string | X | √ |  | 
| region | string | X | √ |  | 
| last_deployment_status_message | string | X | √ |  | 
| last_updated_date | timestamp | X | √ |  | 
| route_settings | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| auto_deploy | bool | X | √ |  | 
| client_certificate_id | string | X | √ |  | 
| deployment_id | string | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| account_id | string | X | √ |  | 


