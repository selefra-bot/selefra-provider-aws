# Table: aws_cloudhsmv2_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tag_list | json | X | √ |  | 
| arn | string | √ | √ |  | 
| cluster_id | string | X | √ |  | 
| pre_co_password | string | X | √ |  | 
| security_group | string | X | √ |  | 
| state | string | X | √ |  | 
| state_message | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| certificates | json | X | √ |  | 
| hsm_type | string | X | √ |  | 
| subnet_mapping | json | X | √ |  | 
| backup_retention_policy | json | X | √ |  | 
| hsms | json | X | √ |  | 
| source_backup_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| backup_policy | string | X | √ |  | 
| create_timestamp | timestamp | X | √ |  | 
| vpc_id | string | X | √ |  | 


