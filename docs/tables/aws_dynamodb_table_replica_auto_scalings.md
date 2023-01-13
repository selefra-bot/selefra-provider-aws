# Table: aws_dynamodb_table_replica_auto_scalings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| global_secondary_indexes | json | X | √ |  | 
| region_name | string | X | √ |  | 
| replica_provisioned_read_capacity_auto_scaling_settings | json | X | √ |  | 
| replica_provisioned_write_capacity_auto_scaling_settings | json | X | √ |  | 
| account_id | string | X | √ |  | 
| table_arn | string | X | √ |  | 
| replica_status | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_dynamodb_tables_selefra_id | string | X | X | fk to aws_dynamodb_tables.selefra_id | 


