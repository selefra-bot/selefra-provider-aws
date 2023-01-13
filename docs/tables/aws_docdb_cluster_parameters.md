# Table: aws_docdb_cluster_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| data_type | string | X | √ |  | 
| allowed_values | string | X | √ |  | 
| apply_method | string | X | √ |  | 
| apply_type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| parameter_name | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 
| source | string | X | √ |  | 
| aws_docdb_engine_versions_selefra_id | string | X | X | fk to aws_docdb_engine_versions.selefra_id | 
| account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| parameter_value | string | X | √ |  | 


