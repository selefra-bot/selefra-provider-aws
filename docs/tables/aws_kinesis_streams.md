# Table: aws_kinesis_streams

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| retention_period_hours | big_int | X | √ |  | 
| stream_creation_timestamp | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| stream_arn | string | X | √ |  | 
| stream_status | string | X | √ |  | 
| consumer_count | big_int | X | √ |  | 
| encryption_type | string | X | √ |  | 
| region | string | X | √ |  | 
| open_shard_count | big_int | X | √ |  | 
| stream_name | string | X | √ |  | 
| stream_mode_details | json | X | √ |  | 
| enhanced_monitoring | json | X | √ |  | 
| key_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


