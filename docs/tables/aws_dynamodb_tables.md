# Table: aws_dynamodb_tables

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| archival_summary | json | X | √ |  | 
| replicas | json | X | √ |  | 
| table_class_summary | json | X | √ |  | 
| tags | json | X | √ |  | 
| provisioned_throughput | json | X | √ |  | 
| table_arn | string | X | √ |  | 
| latest_stream_label | string | X | √ |  | 
| table_id | string | X | √ |  | 
| table_size_bytes | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| key_schema | json | X | √ |  | 
| latest_stream_arn | string | X | √ |  | 
| local_secondary_indexes | json | X | √ |  | 
| stream_specification | json | X | √ |  | 
| attribute_definitions | json | X | √ |  | 
| billing_mode_summary | json | X | √ |  | 
| global_secondary_indexes | json | X | √ |  | 
| sse_description | json | X | √ |  | 
| region | string | X | √ |  | 
| item_count | big_int | X | √ |  | 
| restore_summary | json | X | √ |  | 
| table_name | string | X | √ |  | 
| table_status | string | X | √ |  | 
| creation_date_time | timestamp | X | √ |  | 
| global_table_version | string | X | √ |  | 


