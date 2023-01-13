# Table: aws_rds_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| performance_insights_kms_key_id | string | X | √ |  | 
| activity_stream_status | string | X | √ |  | 
| db_cluster_option_group_memberships | json | X | √ |  | 
| db_system_id | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| performance_insights_enabled | bool | X | √ |  | 
| tags | json | X | √ |  | 
| backtrack_consumed_change_records | big_int | X | √ |  | 
| db_cluster_arn | string | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| db_subnet_group | string | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| engine_mode | string | X | √ |  | 
| status | string | X | √ |  | 
| activity_stream_mode | string | X | √ |  | 
| associated_roles | json | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| tag_list | json | X | √ |  | 
| db_cluster_resource_id | string | X | √ |  | 
| percent_progress | string | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| earliest_backtrack_time | timestamp | X | √ |  | 
| global_write_forwarding_status | string | X | √ |  | 
| activity_stream_kms_key_id | string | X | √ |  | 
| capacity | big_int | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| read_replica_identifiers | string_array | X | √ |  | 
| backtrack_window | big_int | X | √ |  | 
| hosted_zone_id | string | X | √ |  | 
| master_username | string | X | √ |  | 
| scaling_configuration_info | json | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| db_cluster_instance_class | string | X | √ |  | 
| db_cluster_parameter_group | string | X | √ |  | 
| replication_source_identifier | string | X | √ |  | 
| backup_retention_period | big_int | X | √ |  | 
| character_set_name | string | X | √ |  | 
| http_endpoint_enabled | bool | X | √ |  | 
| monitoring_role_arn | string | X | √ |  | 
| reader_endpoint | string | X | √ |  | 
| storage_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| earliest_restorable_time | timestamp | X | √ |  | 
| db_cluster_members | json | X | √ |  | 
| master_user_secret | json | X | √ |  | 
| monitoring_interval | big_int | X | √ |  | 
| arn | string | √ | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| database_name | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| performance_insights_retention_period | big_int | X | √ |  | 
| port | big_int | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| automatic_restart_time | timestamp | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| global_write_forwarding_requested | bool | X | √ |  | 
| multi_az | bool | X | √ |  | 
| region | string | X | √ |  | 
| custom_endpoints | string_array | X | √ |  | 
| iops | big_int | X | √ |  | 
| endpoint | string | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| serverless_v2_scaling_configuration | json | X | √ |  | 
| clone_group_id | string | X | √ |  | 
| domain_memberships | json | X | √ |  | 
| engine | string | X | √ |  | 
| network_type | string | X | √ |  | 
| activity_stream_kinesis_stream_name | string | X | √ |  | 
| cross_account_clone | bool | X | √ |  | 


