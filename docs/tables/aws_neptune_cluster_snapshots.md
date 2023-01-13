# Table: aws_neptune_cluster_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| engine | string | X | √ |  | 
| percent_progress | big_int | X | √ |  | 
| port | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| master_username | string | X | √ |  | 
| source_db_cluster_snapshot_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| db_cluster_snapshot_arn | string | X | √ |  | 
| db_cluster_snapshot_identifier | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| snapshot_type | string | X | √ |  | 
| region | string | X | √ |  | 
| attributes | json | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| license_model | string | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 


