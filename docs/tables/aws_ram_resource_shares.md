# Table: aws_ram_resource_shares

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| owning_account_id | string | X | √ |  | 
| status | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| status_message | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_time | timestamp | X | √ |  | 
| resource_share_arn | string | X | √ |  | 
| allow_external_principals | bool | X | √ |  | 
| feature_set | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


