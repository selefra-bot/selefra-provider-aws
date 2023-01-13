# Table: aws_kafka_nodes

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| cluster_arn | string | X | √ |  | 
| added_to_cluster_time | string | X | √ |  | 
| broker_node_info | json | X | √ |  | 
| node_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| instance_type | string | X | √ |  | 
| node_type | string | X | √ |  | 
| zookeeper_node_info | json | X | √ |  | 
| aws_kafka_clusters_selefra_id | string | X | X | fk to aws_kafka_clusters.selefra_id | 


