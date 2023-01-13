# Table: aws_apigatewayv2_api_deployments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| auto_deployed | bool | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| api_id | string | X | √ |  | 
| deployment_id | string | X | √ |  | 
| deployment_status | string | X | √ |  | 
| deployment_status_message | string | X | √ |  | 
| description | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| account_id | string | X | √ |  | 


