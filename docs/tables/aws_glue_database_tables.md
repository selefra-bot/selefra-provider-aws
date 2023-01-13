# Table: aws_glue_database_tables

## Primary Keys 

```
database_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| database_arn | string | √ | √ |  | 
| database_name | string | X | √ |  | 
| partition_keys | json | X | √ |  | 
| target_table | json | X | √ |  | 
| update_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| description | string | X | √ |  | 
| is_registered_with_lake_formation | bool | X | √ |  | 
| retention | big_int | X | √ |  | 
| storage_descriptor | json | X | √ |  | 
| view_expanded_text | string | X | √ |  | 
| view_original_text | string | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| created_by | string | X | √ |  | 
| last_access_time | timestamp | X | √ |  | 
| last_analyzed_time | timestamp | X | √ |  | 
| owner | string | X | √ |  | 
| table_type | string | X | √ |  | 
| version_id | string | X | √ |  | 
| aws_glue_databases_selefra_id | string | X | X | fk to aws_glue_databases.selefra_id | 
| account_id | string | X | √ |  | 
| name | string | X | √ |  | 
| catalog_id | string | X | √ |  | 
| parameters | json | X | √ |  | 


