# Table: aws_docdb_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| vpc_security_groups | json | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| status_infos | json | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| db_instance_identifier | string | X | √ |  | 
| engine | string | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| endpoint | json | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| db_instance_status | string | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| region | string | X | √ |  | 
| ca_certificate_identifier | string | X | √ |  | 
| db_instance_class | string | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| engine_version | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| aws_docdb_clusters_selefra_id | string | X | X | fk to aws_docdb_clusters.selefra_id | 
| account_id | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| backup_retention_period | big_int | X | √ |  | 
| dbi_resource_id | string | X | √ |  | 
| db_subnet_group | json | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| promotion_tier | big_int | X | √ |  | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| db_instance_arn | string | X | √ |  | 


