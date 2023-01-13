# Table: aws_elasticache_replication_groups

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| replication_group_create_time | timestamp | X | √ |  | 
| user_group_ids | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| automatic_failover | string | X | √ |  | 
| description | string | X | √ |  | 
| node_groups | json | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| snapshot_window | string | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| global_replication_group_info | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| log_delivery_configurations | json | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| network_type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| at_rest_encryption_enabled | bool | X | √ |  | 
| auth_token_last_modified_date | timestamp | X | √ |  | 
| cluster_enabled | bool | X | √ |  | 
| configuration_endpoint | json | X | √ |  | 
| ip_discovery | string | X | √ |  | 
| member_clusters | string_array | X | √ |  | 
| region | string | X | √ |  | 
| member_clusters_outpost_arns | string_array | X | √ |  | 
| status | string | X | √ |  | 
| transit_encryption_enabled | bool | X | √ |  | 
| arn | string | X | √ |  | 
| auth_token_enabled | bool | X | √ |  | 
| replication_group_id | string | X | √ |  | 
| transit_encryption_mode | string | X | √ |  | 
| data_tiering | string | X | √ |  | 
| multi_az | string | X | √ |  | 
| snapshot_retention_limit | big_int | X | √ |  | 
| snapshotting_cluster_id | string | X | √ |  | 


