# Table: aws_s3_accounts

## Primary Keys 

```
account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | √ | √ |  | 
| public_access_block_configuration | json | X | √ |  | 
| config_exists | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


