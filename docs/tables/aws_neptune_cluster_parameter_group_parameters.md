# Table: aws_neptune_cluster_parameter_group_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| minimum_engine_version | string | X | √ |  | 
| parameter_name | string | X | √ |  | 
| source | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_neptune_cluster_parameter_groups_selefra_id | string | X | X | fk to aws_neptune_cluster_parameter_groups.selefra_id | 
| cluster_parameter_group_arn | string | X | √ |  | 
| allowed_values | string | X | √ |  | 
| description | string | X | √ |  | 
| parameter_value | string | X | √ |  | 
| region | string | X | √ |  | 
| apply_method | string | X | √ |  | 
| data_type | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| apply_type | string | X | √ |  | 


