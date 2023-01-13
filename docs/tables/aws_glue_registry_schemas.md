# Table: aws_glue_registry_schemas

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| data_format | string | X | √ |  | 
| latest_schema_version | big_int | X | √ |  | 
| next_schema_version | big_int | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| registry_arn | string | X | √ |  | 
| schema_arn | string | X | √ |  | 
| schema_status | string | X | √ |  | 
| aws_glue_registries_selefra_id | string | X | X | fk to aws_glue_registries.selefra_id | 
| tags | json | X | √ |  | 
| compatibility | string | X | √ |  | 
| registry_name | string | X | √ |  | 
| schema_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| created_time | string | X | √ |  | 
| updated_time | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| description | string | X | √ |  | 
| schema_checkpoint | big_int | X | √ |  | 


