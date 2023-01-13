# Table: aws_neptune_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cross_account_clone | bool | X | √ |  | 
| endpoint | string | X | √ |  | 
| engine | string | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| master_username | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| backup_retention_period | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| clone_group_id | string | X | √ |  | 
| database_name | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| percent_progress | string | X | √ |  | 
| read_replica_identifiers | string_array | X | √ |  | 
| region | string | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| earliest_restorable_time | timestamp | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| serverless_v2_scaling_configuration | json | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| db_cluster_resource_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| hosted_zone_id | string | X | √ |  | 
| replication_source_identifier | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| db_cluster_identifier | string | X | √ |  | 
| db_cluster_parameter_group | string | X | √ |  | 
| db_cluster_members | json | X | √ |  | 
| automatic_restart_time | timestamp | X | √ |  | 
| db_cluster_option_group_memberships | json | X | √ |  | 
| db_subnet_group | string | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| character_set_name | string | X | √ |  | 
| db_cluster_arn | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| port | big_int | X | √ |  | 
| reader_endpoint | string | X | √ |  | 
| associated_roles | json | X | √ |  | 


