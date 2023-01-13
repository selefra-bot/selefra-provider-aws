# Table: aws_fsx_data_repository_associations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| failure_details | json | X | √ |  | 
| resource_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| s3 | json | X | √ |  | 
| arn | string | √ | √ |  | 
| batch_import_meta_data_on_create | bool | X | √ |  | 
| data_repository_path | string | X | √ |  | 
| data_repository_subdirectories | string_array | X | √ |  | 
| nfs | json | X | √ |  | 
| association_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| file_system_path | string | X | √ |  | 
| imported_file_chunk_size | big_int | X | √ |  | 
| file_system_id | string | X | √ |  | 
| lifecycle | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| file_cache_id | string | X | √ |  | 
| file_cache_path | string | X | √ |  | 


