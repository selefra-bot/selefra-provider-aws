# Table: aws_lambda_layer_version_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| revision_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| layer_version | int | X | √ |  | 
| policy | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| aws_lambda_layer_versions_selefra_id | string | X | X | fk to aws_lambda_layer_versions.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| layer_version_arn | string | X | √ |  | 


