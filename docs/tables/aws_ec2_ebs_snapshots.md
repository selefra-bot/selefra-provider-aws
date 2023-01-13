# Table: aws_ec2_ebs_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| owner_id | string | X | √ |  | 
| restore_expiry_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| attribute | json | X | √ |  | 
| tags | json | X | √ |  | 
| data_encryption_key_id | string | X | √ |  | 
| description | string | X | √ |  | 
| progress | string | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| volume_id | string | X | √ |  | 
| volume_size | big_int | X | √ |  | 
| arn | string | √ | √ |  | 
| encrypted | bool | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| state | string | X | √ |  | 
| storage_tier | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| owner_alias | string | X | √ |  | 
| snapshot_id | string | X | √ |  | 
| state_message | string | X | √ |  | 


