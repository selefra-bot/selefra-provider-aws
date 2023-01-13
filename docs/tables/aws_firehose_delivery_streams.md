# Table: aws_firehose_delivery_streams

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| has_more_destinations | bool | X | √ |  | 
| create_timestamp | timestamp | X | √ |  | 
| failure_description | json | X | √ |  | 
| delivery_stream_status | string | X | √ |  | 
| delivery_stream_type | string | X | √ |  | 
| version_id | string | X | √ |  | 
| delivery_stream_name | string | X | √ |  | 
| destinations | json | X | √ |  | 
| source | json | X | √ |  | 
| account_id | string | X | √ |  | 
| delivery_stream_arn | string | X | √ |  | 
| delivery_stream_encryption_configuration | json | X | √ |  | 
| last_update_timestamp | timestamp | X | √ |  | 


