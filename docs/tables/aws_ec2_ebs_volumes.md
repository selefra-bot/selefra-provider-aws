# Table: aws_ec2_ebs_volumes

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| availability_zone | string | X | √ |  | 
| iops | big_int | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| multi_attach_enabled | bool | X | √ |  | 
| size | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| outpost_arn | string | X | √ |  | 
| throughput | big_int | X | √ |  | 
| volume_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| attachments | json | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| snapshot_id | string | X | √ |  | 
| state | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| encrypted | bool | X | √ |  | 
| fast_restored | bool | X | √ |  | 
| volume_type | string | X | √ |  | 


