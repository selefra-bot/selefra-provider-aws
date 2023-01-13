# Table: aws_kafka_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cluster_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| active_operation_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| cluster_name | string | X | √ |  | 
| current_version | string | X | √ |  | 
| state | string | X | √ |  | 
| arn | string | √ | √ |  | 
| serverless | json | X | √ |  | 
| provisioned | json | X | √ |  | 
| state_info | json | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 


