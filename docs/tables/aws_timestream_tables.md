# Table: aws_timestream_tables

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| table_name | string | X | √ |  | 
| table_status | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| magnetic_store_write_properties | json | X | √ |  | 
| aws_timestream_databases_selefra_id | string | X | X | fk to aws_timestream_databases.selefra_id | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| database_name | string | X | √ |  | 
| retention_properties | json | X | √ |  | 


