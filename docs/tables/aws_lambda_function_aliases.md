# Table: aws_lambda_function_aliases

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| function_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 
| alias_configuration | json | X | √ |  | 
| url_config | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


