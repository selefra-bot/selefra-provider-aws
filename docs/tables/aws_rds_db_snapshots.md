# Table: aws_rds_db_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| engine | string | X | √ |  | 
| percent_progress | big_int | X | √ |  | 
| storage_throughput | big_int | X | √ |  | 
| vpc_id | string | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| timezone | string | X | √ |  | 
| region | string | X | √ |  | 
| attributes | json | X | √ |  | 
| availability_zone | string | X | √ |  | 
| storage_type | string | X | √ |  | 
| tags | json | X | √ |  | 
| encrypted | bool | X | √ |  | 
| iops | big_int | X | √ |  | 
| source_db_snapshot_identifier | string | X | √ |  | 
| source_region | string | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| port | big_int | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| snapshot_target | string | X | √ |  | 
| option_group_name | string | X | √ |  | 
| snapshot_database_time | timestamp | X | √ |  | 
| db_instance_identifier | string | X | √ |  | 
| dbi_resource_id | string | X | √ |  | 
| license_model | string | X | √ |  | 
| master_username | string | X | √ |  | 
| status | string | X | √ |  | 
| tde_credential_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| db_snapshot_arn | string | X | √ |  | 
| original_snapshot_create_time | timestamp | X | √ |  | 
| processor_features | json | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| tag_list | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| allocated_storage | big_int | X | √ |  | 
| db_snapshot_identifier | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 


