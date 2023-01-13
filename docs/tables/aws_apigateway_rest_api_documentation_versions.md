# Table: aws_apigateway_rest_api_documentation_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| description | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| rest_api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| version | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 


