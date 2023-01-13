# Table: aws_eventbridge_archives

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state_reason | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| archive_name | string | X | √ |  | 
| retention_days | big_int | X | √ |  | 
| size_bytes | big_int | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| event_count | big_int | X | √ |  | 
| event_source_arn | string | X | √ |  | 


