# Table: aws_rds_cluster_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| source | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| data_type | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| account_id | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 
| apply_type | string | X | √ |  | 
| description | string | X | √ |  | 
| parameter_name | string | X | √ |  | 
| supported_engine_modes | string_array | X | √ |  | 
| aws_rds_engine_versions_selefra_id | string | X | X | fk to aws_rds_engine_versions.selefra_id | 
| allowed_values | string | X | √ |  | 
| apply_method | string | X | √ |  | 
| parameter_value | string | X | √ |  | 


