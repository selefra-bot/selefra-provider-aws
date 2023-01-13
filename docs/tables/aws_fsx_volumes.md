# Table: aws_fsx_volumes

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| administrative_actions | json | X | √ |  | 
| ontap_configuration | json | X | √ |  | 
| open_zfs_configuration | json | X | √ |  | 
| lifecycle_transition_reason | json | X | √ |  | 
| volume_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| lifecycle | string | X | √ |  | 
| tags | json | X | √ |  | 
| name | string | X | √ |  | 
| volume_id | string | X | √ |  | 
| resource_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| file_system_id | string | X | √ |  | 


