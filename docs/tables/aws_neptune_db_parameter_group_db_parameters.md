# Table: aws_neptune_db_parameter_group_db_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| data_type | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| source | string | X | √ |  | 
| account_id | string | X | √ |  | 
| apply_method | string | X | √ |  | 
| description | string | X | √ |  | 
| parameter_name | string | X | √ |  | 
| parameter_value | string | X | √ |  | 
| aws_neptune_db_parameter_groups_selefra_id | string | X | X | fk to aws_neptune_db_parameter_groups.selefra_id | 
| allowed_values | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| apply_type | string | X | √ |  | 
| db_parameter_group_arn | string | X | √ |  | 
| region | string | X | √ |  | 


