# Table: aws_elasticache_clusters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| at_rest_encryption_enabled | bool | X | √ |  | 
| engine | string | X | √ |  | 
| notification_configuration | json | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| preferred_availability_zone | string | X | √ |  | 
| cache_cluster_create_time | timestamp | X | √ |  | 
| cache_cluster_id | string | X | √ |  | 
| cache_cluster_status | string | X | √ |  | 
| network_type | string | X | √ |  | 
| log_delivery_configurations | json | X | √ |  | 
| snapshot_retention_limit | big_int | X | √ |  | 
| snapshot_window | string | X | √ |  | 
| transit_encryption_enabled | bool | X | √ |  | 
| auth_token_enabled | bool | X | √ |  | 
| cache_security_groups | json | X | √ |  | 
| cache_subnet_group_name | string | X | √ |  | 
| num_cache_nodes | big_int | X | √ |  | 
| preferred_outpost_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| cache_parameter_group | json | X | √ |  | 
| engine_version | string | X | √ |  | 
| transit_encryption_mode | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| cache_nodes | json | X | √ |  | 
| ip_discovery | string | X | √ |  | 
| security_groups | json | X | √ |  | 
| auth_token_last_modified_date | timestamp | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| replication_group_log_delivery_enabled | bool | X | √ |  | 
| arn | string | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| client_download_landing_page | string | X | √ |  | 
| configuration_endpoint | json | X | √ |  | 
| replication_group_id | string | X | √ |  | 


