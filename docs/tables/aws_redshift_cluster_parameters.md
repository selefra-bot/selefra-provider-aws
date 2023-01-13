# Table: aws_redshift_cluster_parameters

## Primary Keys 

```
cluster_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| parameter_value | string | X | √ |  | 
| allowed_values | string | X | √ |  | 
| region | string | X | √ |  | 
| apply_type | string | X | √ |  | 
| data_type | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 
| aws_redshift_cluster_parameter_groups_selefra_id | string | X | X | fk to aws_redshift_cluster_parameter_groups.selefra_id | 
| account_id | string | X | √ |  | 
| cluster_arn | string | √ | √ | `The Amazon Resource Name (ARN) for the resource.` | 
| parameter_name | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| source | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


