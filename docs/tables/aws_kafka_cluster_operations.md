# Table: aws_kafka_cluster_operations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| operation_arn | string | X | √ |  | 
| operation_state | string | X | √ |  | 
| target_cluster_info | json | X | √ |  | 
| tags | json | X | √ |  | 
| end_time | timestamp | X | √ |  | 
| operation_steps | json | X | √ |  | 
| source_cluster_info | json | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| client_request_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| operation_type | string | X | √ |  | 
| aws_kafka_clusters_selefra_id | string | X | X | fk to aws_kafka_clusters.selefra_id | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| error_info | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


