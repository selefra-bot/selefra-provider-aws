# Table: aws_docdb_orderable_db_instance_options

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| engine_version | string | X | √ |  | 
| license_model | string | X | √ |  | 
| vpc | bool | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_docdb_engine_versions_selefra_id | string | X | X | fk to aws_docdb_engine_versions.selefra_id | 
| availability_zones | json | X | √ |  | 
| db_instance_class | string | X | √ |  | 
| engine | string | X | √ |  | 


