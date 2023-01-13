# Table: aws_backup_vaults

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| backup_vault_arn | string | X | √ |  | 
| access_policy | json | X | √ |  | 
| creator_request_id | string | X | √ |  | 
| lock_date | timestamp | X | √ |  | 
| locked | bool | X | √ |  | 
| min_retention_days | big_int | X | √ |  | 
| number_of_recovery_points | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| backup_vault_name | string | X | √ |  | 
| encryption_key_arn | string | X | √ |  | 
| max_retention_days | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| notifications | json | X | √ |  | 
| creation_date | timestamp | X | √ |  | 


