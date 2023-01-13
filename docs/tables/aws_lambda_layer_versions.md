# Table: aws_lambda_layer_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| license_info | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| layer_arn | string | X | √ |  | 
| compatible_runtimes | string_array | X | √ |  | 
| layer_version_arn | string | X | √ |  | 
| version | big_int | X | √ |  | 
| aws_lambda_layers_selefra_id | string | X | X | fk to aws_lambda_layers.selefra_id | 
| region | string | X | √ |  | 
| compatible_architectures | string_array | X | √ |  | 
| created_date | string | X | √ |  | 
| description | string | X | √ |  | 


