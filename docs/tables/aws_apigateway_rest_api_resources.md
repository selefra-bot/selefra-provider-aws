# Table: aws_apigateway_rest_api_resources

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| parent_id | string | X | √ |  | 
| path_part | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| id | string | X | √ |  | 
| path | string | X | √ |  | 
| resource_methods | json | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| selefra_id | string | √ | √ | random id | 


