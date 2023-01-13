# Table: aws_apigatewayv2_api_models

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| name | string | X | √ |  | 
| content_type | string | X | √ |  | 
| description | string | X | √ |  | 
| schema | string | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| api_arn | string | X | √ |  | 
| api_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| model_template | string | X | √ |  | 
| model_id | string | X | √ |  | 


