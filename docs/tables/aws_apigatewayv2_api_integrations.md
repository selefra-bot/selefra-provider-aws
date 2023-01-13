# Table: aws_apigatewayv2_api_integrations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| region | string | X | √ |  | 
| credentials_arn | string | X | √ |  | 
| integration_response_selection_expression | string | X | √ |  | 
| integration_type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| integration_id | string | X | √ |  | 
| integration_method | string | X | √ |  | 
| integration_uri | string | X | √ |  | 
| passthrough_behavior | string | X | √ |  | 
| payload_format_version | string | X | √ |  | 
| request_parameters | json | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| timeout_in_millis | big_int | X | √ |  | 
| description | string | X | √ |  | 
| response_parameters | json | X | √ |  | 
| content_handling_strategy | string | X | √ |  | 
| api_gateway_managed | bool | X | √ |  | 
| connection_type | string | X | √ |  | 
| api_id | string | X | √ |  | 
| integration_subtype | string | X | √ |  | 
| template_selection_expression | string | X | √ |  | 
| connection_id | string | X | √ |  | 
| request_templates | json | X | √ |  | 
| api_arn | string | X | √ |  | 
| tls_config | json | X | √ |  | 


