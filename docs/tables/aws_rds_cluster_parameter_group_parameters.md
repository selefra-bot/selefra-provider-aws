# Table: aws_rds_cluster_parameter_group_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| apply_type | string | X | √ |  | 
| data_type | string | X | √ |  | 
| description | string | X | √ |  | 
| aws_rds_cluster_parameter_groups_selefra_id | string | X | X | fk to aws_rds_cluster_parameter_groups.selefra_id | 
| is_modifiable | bool | X | √ |  | 
| parameter_value | string | X | √ |  | 
| supported_engine_modes | string_array | X | √ |  | 
| cluster_parameter_group_arn | string | X | √ |  | 
| allowed_values | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| source | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| apply_method | string | X | √ |  | 
| parameter_name | string | X | √ |  | 


