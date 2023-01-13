# Table: aws_rds_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| dbi_resource_id | string | X | √ |  | 
| monitoring_role_arn | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| availability_zone | string | X | √ |  | 
| db_subnet_group | json | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| character_set_name | string | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| master_username | string | X | √ |  | 
| storage_throughput | big_int | X | √ |  | 
| automation_mode | string | X | √ |  | 
| backup_retention_period | big_int | X | √ |  | 
| db_security_groups | json | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| tde_credential_arn | string | X | √ |  | 
| master_user_secret | json | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| processor_features | json | X | √ |  | 
| custom_iam_instance_profile | string | X | √ |  | 
| db_instance_identifier | string | X | √ |  | 
| db_instance_port | big_int | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| license_model | string | X | √ |  | 
| read_replica_db_cluster_identifiers | string_array | X | √ |  | 
| enhanced_monitoring_resource_arn | string | X | √ |  | 
| network_type | string | X | √ |  | 
| replica_mode | string | X | √ |  | 
| storage_type | string | X | √ |  | 
| performance_insights_enabled | bool | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| aws_backup_recovery_point_arn | string | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| db_parameter_groups | json | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| nchar_character_set_name | string | X | √ |  | 
| read_replica_db_instance_identifiers | string_array | X | √ |  | 
| automatic_restart_time | timestamp | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| domain_memberships | json | X | √ |  | 
| engine | string | X | √ |  | 
| performance_insights_retention_period | big_int | X | √ |  | 
| resume_full_automation_mode_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| activity_stream_status | string | X | √ |  | 
| db_system_id | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| promotion_tier | big_int | X | √ |  | 
| secondary_availability_zone | string | X | √ |  | 
| status_infos | json | X | √ |  | 
| activity_stream_kms_key_id | string | X | √ |  | 
| endpoint | json | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| tag_list | json | X | √ |  | 
| activity_stream_kinesis_stream_name | string | X | √ |  | 
| db_instance_status | string | X | √ |  | 
| monitoring_interval | big_int | X | √ |  | 
| option_group_memberships | json | X | √ |  | 
| arn | string | √ | √ |  | 
| activity_stream_policy_status | string | X | √ |  | 
| backup_target | string | X | √ |  | 
| certificate_details | json | X | √ |  | 
| max_allocated_storage | big_int | X | √ |  | 
| activity_stream_engine_native_audit_fields_included | bool | X | √ |  | 
| ca_certificate_identifier | string | X | √ |  | 
| db_instance_automated_backups_replications | json | X | √ |  | 
| timezone | string | X | √ |  | 
| account_id | string | X | √ |  | 
| db_instance_arn | string | X | √ |  | 
| listener_endpoint | json | X | √ |  | 
| performance_insights_kms_key_id | string | X | √ |  | 
| read_replica_source_db_instance_identifier | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| customer_owned_ip_enabled | bool | X | √ |  | 
| engine_version | string | X | √ |  | 
| iops | big_int | X | √ |  | 
| activity_stream_mode | string | X | √ |  | 
| associated_roles | json | X | √ |  | 
| db_instance_class | string | X | √ |  | 
| db_name | string | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 


