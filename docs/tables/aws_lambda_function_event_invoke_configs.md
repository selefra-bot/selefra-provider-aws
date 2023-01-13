# Table: aws_lambda_function_event_invoke_configs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| function_arn | string | X | √ |  | 
| maximum_event_age_in_seconds | big_int | X | √ |  | 
| maximum_retry_attempts | big_int | X | √ |  | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| account_id | string | X | √ |  | 
| destination_config | json | X | √ |  | 
| last_modified | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


