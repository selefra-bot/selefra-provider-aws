# Table: aws_apigateway_usage_plan_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_apigateway_usage_plans_selefra_id | string | X | X | fk to aws_apigateway_usage_plans.selefra_id | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| id | string | X | √ |  | 
| type | string | X | √ |  | 
| value | string | X | √ |  | 
| account_id | string | X | √ |  | 
| usage_plan_arn | string | X | √ |  | 


