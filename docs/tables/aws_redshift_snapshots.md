# Table: aws_redshift_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| actual_incremental_backup_size_in_mega_bytes | float | X | √ |  | 
| db_name | string | X | √ |  | 
| cluster_identifier | string | X | √ |  | 
| current_backup_rate_in_mega_bytes_per_second | float | X | √ |  | 
| estimated_seconds_to_completion | big_int | X | √ |  | 
| aws_redshift_clusters_selefra_id | string | X | X | fk to aws_redshift_clusters.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| total_backup_size_in_mega_bytes | float | X | √ |  | 
| engine_full_version | string | X | √ |  | 
| enhanced_vpc_routing | bool | X | √ |  | 
| manual_snapshot_retention_period | big_int | X | √ |  | 
| node_type | string | X | √ |  | 
| snapshot_identifier | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| availability_zone | string | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| maintenance_track_name | string | X | √ |  | 
| restorable_node_types | string_array | X | √ |  | 
| snapshot_retention_start_time | timestamp | X | √ |  | 
| accounts_with_restore_access | json | X | √ |  | 
| elapsed_time_in_seconds | big_int | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| master_username | string | X | √ |  | 
| number_of_nodes | big_int | X | √ |  | 
| cluster_version | string | X | √ |  | 
| encrypted | bool | X | √ |  | 
| status | string | X | √ |  | 
| port | big_int | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| owner_account | string | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| source_region | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ | `ARN of the snapshot.` | 
| backup_progress_in_mega_bytes | float | X | √ |  | 
| encrypted_with_hsm | bool | X | √ |  | 
| manual_snapshot_remaining_days | big_int | X | √ |  | 


