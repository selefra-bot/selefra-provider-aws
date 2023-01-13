# Table: aws_neptune_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| read_replica_source_db_instance_identifier | string | X | √ |  | 
| ca_certificate_identifier | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| status_infos | json | X | √ |  | 
| tde_credential_arn | string | X | √ |  | 
| read_replica_db_cluster_identifiers | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| db_instance_identifier | string | X | √ |  | 
| db_security_groups | json | X | √ |  | 
| db_instance_port | big_int | X | √ |  | 
| engine | string | X | √ |  | 
| backup_retention_period | big_int | X | √ |  | 
| db_name | string | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| read_replica_db_instance_identifiers | string_array | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| character_set_name | string | X | √ |  | 
| endpoint | json | X | √ |  | 
| enhanced_monitoring_resource_arn | string | X | √ |  | 
| iops | big_int | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| db_instance_class | string | X | √ |  | 
| region | string | X | √ |  | 
| db_instance_arn | string | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| option_group_memberships | json | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| db_parameter_groups | json | X | √ |  | 
| db_subnet_group | json | X | √ |  | 
| performance_insights_kms_key_id | string | X | √ |  | 
| timezone | string | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| domain_memberships | json | X | √ |  | 
| master_username | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| secondary_availability_zone | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| engine_version | string | X | √ |  | 
| license_model | string | X | √ |  | 
| storage_type | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| db_instance_status | string | X | √ |  | 
| dbi_resource_id | string | X | √ |  | 
| monitoring_role_arn | string | X | √ |  | 
| promotion_tier | big_int | X | √ |  | 
| monitoring_interval | big_int | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| performance_insights_enabled | bool | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 


