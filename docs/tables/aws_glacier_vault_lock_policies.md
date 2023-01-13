# Table: aws_glacier_vault_lock_policies

## Primary Keys 

```
vault_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| vault_arn | string | √ | √ |  | 
| creation_date | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| policy | string | X | √ |  | 
| expiration_date | string | X | √ |  | 
| state | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| aws_glacier_vaults_selefra_id | string | X | X | fk to aws_glacier_vaults.selefra_id | 


