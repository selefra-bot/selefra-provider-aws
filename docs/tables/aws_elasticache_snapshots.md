# Table: aws_elasticache_snapshots

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| replication_group_description | string | X | √ |  | 
| replication_group_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| num_cache_nodes | big_int | X | √ |  | 
| snapshot_status | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| region | string | X | √ |  | 
| snapshot_window | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| automatic_failover | string | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| cache_subnet_group_name | string | X | √ |  | 
| engine | string | X | √ |  | 
| preferred_availability_zone | string | X | √ |  | 
| cache_cluster_id | string | X | √ |  | 
| data_tiering | string | X | √ |  | 
| node_snapshots | json | X | √ |  | 
| num_node_groups | big_int | X | √ |  | 
| port | big_int | X | √ |  | 
| preferred_outpost_arn | string | X | √ |  | 
| snapshot_retention_limit | big_int | X | √ |  | 
| cache_cluster_create_time | timestamp | X | √ |  | 
| cache_parameter_group_name | string | X | √ |  | 
| snapshot_name | string | X | √ |  | 
| snapshot_source | string | X | √ |  | 
| topic_arn | string | X | √ |  | 


