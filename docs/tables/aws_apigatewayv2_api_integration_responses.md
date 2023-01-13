# Table: aws_apigatewayv2_api_integration_responses

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| integration_response_key | string | X | √ |  | 
| content_handling_strategy | string | X | √ |  | 
| integration_response_id | string | X | √ |  | 
| response_templates | json | X | √ |  | 
| aws_apigatewayv2_api_integrations_selefra_id | string | X | X | fk to aws_apigatewayv2_api_integrations.selefra_id | 
| region | string | X | √ |  | 
| api_integration_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| template_selection_expression | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| integration_id | string | X | √ |  | 
| response_parameters | json | X | √ |  | 


