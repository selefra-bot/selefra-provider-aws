# Table: aws_apigateway_rest_api_deployments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| id | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| region | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| api_summary | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


