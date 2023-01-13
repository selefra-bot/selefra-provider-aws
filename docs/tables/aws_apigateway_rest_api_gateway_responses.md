# Table: aws_apigateway_rest_api_gateway_responses

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| response_parameters | json | X | √ |  | 
| response_type | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| default_response | bool | X | √ |  | 
| response_templates | json | X | √ |  | 
| status_code | string | X | √ |  | 


