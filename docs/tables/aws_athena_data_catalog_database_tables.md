# Table: aws_athena_data_catalog_database_tables

## Primary Keys 

```
data_catalog_arn, data_catalog_database_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| columns | json | X | √ |  | 
| parameters | json | X | √ |  | 
| partition_keys | json | X | √ |  | 
| aws_athena_data_catalog_databases_selefra_id | string | X | X | fk to aws_athena_data_catalog_databases.selefra_id | 
| data_catalog_arn | string | X | √ |  | 
| data_catalog_database_name | string | X | √ |  | 
| name | string | X | √ |  | 
| last_access_time | timestamp | X | √ |  | 
| table_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| create_time | timestamp | X | √ |  | 


