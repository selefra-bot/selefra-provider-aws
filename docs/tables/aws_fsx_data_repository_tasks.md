# Table: aws_fsx_data_repository_tasks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_time | timestamp | X | √ |  | 
| type | string | X | √ |  | 
| end_time | timestamp | X | √ |  | 
| failure_details | json | X | √ |  | 
| status | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| task_id | string | X | √ |  | 
| paths | string_array | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| capacity_to_release | big_int | X | √ |  | 
| file_cache_id | string | X | √ |  | 
| file_system_id | string | X | √ |  | 
| report | json | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| lifecycle | string | X | √ |  | 


