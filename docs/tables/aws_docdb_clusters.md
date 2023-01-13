# Table: aws_docdb_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| multi_az | bool | X | √ |  | 
| reader_endpoint | string | X | √ |  | 
| region | string | X | √ |  | 
| backup_retention_period | big_int | X | √ |  | 
| engine | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| arn | string | √ | √ |  | 
| availability_zones | string_array | X | √ |  | 
| replication_source_identifier | string | X | √ |  | 
| db_cluster_arn | string | X | √ |  | 
| hosted_zone_id | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| status | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| port | big_int | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| read_replica_identifiers | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| master_username | string | X | √ |  | 
| percent_progress | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| clone_group_id | string | X | √ |  | 
| db_cluster_members | json | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| endpoint | string | X | √ |  | 
| tags | json | X | √ |  | 
| associated_roles | json | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| db_subnet_group | string | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| cluster_create_time | timestamp | X | √ |  | 
| db_cluster_parameter_group | string | X | √ |  | 
| db_cluster_resource_id | string | X | √ |  | 
| earliest_restorable_time | timestamp | X | √ |  | 


