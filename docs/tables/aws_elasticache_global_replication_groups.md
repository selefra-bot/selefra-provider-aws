# Table: aws_elasticache_global_replication_groups

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| global_replication_group_id | string | X | √ |  | 
| members | json | X | √ |  | 
| transit_encryption_enabled | bool | X | √ |  | 
| region | string | X | √ |  | 
| auth_token_enabled | bool | X | √ |  | 
| engine_version | string | X | √ |  | 
| global_replication_group_description | string | X | √ |  | 
| arn | string | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| engine | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| at_rest_encryption_enabled | bool | X | √ |  | 
| cluster_enabled | bool | X | √ |  | 
| global_node_groups | json | X | √ |  | 


