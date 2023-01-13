# Table: aws_quicksight_ingestions

## Primary Keys 

```
data_set_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| error_info | json | X | √ |  | 
| ingestion_id | string | X | √ |  | 
| ingestion_size_in_bytes | big_int | X | √ |  | 
| row_info | json | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| ingestion_status | string | X | √ |  | 
| request_source | string | X | √ |  | 
| request_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| ingestion_time_in_seconds | big_int | X | √ |  | 
| queue_info | json | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| data_set_arn | string | √ | √ |  | 
| aws_quicksight_data_sets_selefra_id | string | X | X | fk to aws_quicksight_data_sets.selefra_id | 


