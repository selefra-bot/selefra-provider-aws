# Table: aws_apigateway_rest_api_documentation_parts

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| location | json | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| account_id | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| id | string | X | √ |  | 
| properties | string | X | √ |  | 


