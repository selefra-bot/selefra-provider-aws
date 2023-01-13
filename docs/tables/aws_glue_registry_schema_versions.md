# Table: aws_glue_registry_schema_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| registry_schema_arn | string | X | √ |  | 
| created_time | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| schema_arn | string | X | √ |  | 
| version_number | big_int | X | √ |  | 
| aws_glue_registry_schemas_selefra_id | string | X | X | fk to aws_glue_registry_schemas.selefra_id | 
| region | string | X | √ |  | 
| metadata | json | X | √ |  | 
| data_format | string | X | √ |  | 
| schema_definition | string | X | √ |  | 
| schema_version_id | string | X | √ |  | 
| status | string | X | √ |  | 
| result_metadata | json | X | √ |  | 


