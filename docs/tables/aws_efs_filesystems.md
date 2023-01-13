# Table: aws_efs_filesystems

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_time | timestamp | X | √ |  | 
| file_system_id | string | X | √ |  | 
| performance_mode | string | X | √ |  | 
| encrypted | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| throughput_mode | string | X | √ |  | 
| backup_policy_status | string | X | √ |  | 
| life_cycle_state | string | X | √ |  | 
| number_of_mount_targets | big_int | X | √ |  | 
| size_in_bytes | json | X | √ |  | 
| availability_zone_id | string | X | √ |  | 
| availability_zone_name | string | X | √ |  | 
| file_system_arn | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_token | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| name | string | X | √ |  | 
| provisioned_throughput_in_mibps | float | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


