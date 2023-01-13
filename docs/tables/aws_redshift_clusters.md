# Table: aws_redshift_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cluster_parameter_groups | json | X | √ |  | 
| cluster_subnet_group_name | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| cluster_snapshot_copy_status | json | X | √ |  | 
| iam_roles | json | X | √ |  | 
| account_id | string | X | √ |  | 
| cluster_availability_status | string | X | √ |  | 
| cluster_nodes | json | X | √ |  | 
| cluster_version | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| elastic_resize_number_of_node_options | string | X | √ |  | 
| expected_next_snapshot_schedule_time | timestamp | X | √ |  | 
| manual_snapshot_retention_period | big_int | X | √ |  | 
| total_storage_capacity_in_mega_bytes | big_int | X | √ |  | 
| snapshot_schedule_state | string | X | √ |  | 
| tags | json | X | √ |  | 
| encrypted | bool | X | √ |  | 
| expected_next_snapshot_schedule_time_status | string | X | √ |  | 
| hsm_status | json | X | √ |  | 
| snapshot_schedule_identifier | string | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| region | string | X | √ |  | 
| cluster_namespace_arn | string | X | √ |  | 
| elastic_ip_status | json | X | √ |  | 
| endpoint | json | X | √ |  | 
| master_username | string | X | √ |  | 
| cluster_security_groups | json | X | √ |  | 
| maintenance_track_name | string | X | √ |  | 
| modify_status | string | X | √ |  | 
| reserved_node_exchange_status | json | X | √ |  | 
| resize_info | json | X | √ |  | 
| arn | string | √ | √ | `The Amazon Resource Name (ARN) for the resource.` | 
| aqua_configuration | json | X | √ |  | 
| automated_snapshot_retention_period | big_int | X | √ |  | 
| cluster_status | string | X | √ |  | 
| allow_version_upgrade | bool | X | √ |  | 
| data_transfer_progress | json | X | √ |  | 
| deferred_maintenance_windows | json | X | √ |  | 
| pending_actions | string_array | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| cluster_public_key | string | X | √ |  | 
| cluster_revision_number | string | X | √ |  | 
| restore_status | json | X | √ |  | 
| next_maintenance_window_start_time | timestamp | X | √ |  | 
| node_type | string | X | √ |  | 
| logging_status | json | X | √ | `Describes the status of logging for a cluster.` | 
| availability_zone_relocation_status | string | X | √ |  | 
| default_iam_role_arn | string | X | √ |  | 
| enhanced_vpc_routing | bool | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| cluster_identifier | string | X | √ |  | 
| number_of_nodes | big_int | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| db_name | string | X | √ |  | 


