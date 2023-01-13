# Table: aws_rds_cluster_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| engine_mode | string | X | √ |  | 
| tag_list | json | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| attributes | json | X | √ |  | 
| db_system_id | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| db_cluster_snapshot_identifier | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| license_model | string | X | √ |  | 
| source_db_cluster_snapshot_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| engine | string | X | √ |  | 
| port | big_int | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ |  | 
| engine_version | string | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| master_username | string | X | √ |  | 
| percent_progress | big_int | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| arn | string | √ | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| db_cluster_snapshot_arn | string | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 


