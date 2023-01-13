# Table: aws_cloudhsmv2_backups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| backup_id | string | X | √ |  | 
| cluster_id | string | X | √ |  | 
| copy_timestamp | timestamp | X | √ |  | 
| source_cluster | string | X | √ |  | 
| tag_list | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| backup_state | string | X | √ |  | 
| account_id | string | X | √ |  | 
| create_timestamp | timestamp | X | √ |  | 
| source_backup | string | X | √ |  | 
| source_region | string | X | √ |  | 
| delete_timestamp | timestamp | X | √ |  | 
| never_expires | bool | X | √ |  | 


