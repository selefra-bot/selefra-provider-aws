# Table: aws_kms_aliases

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| alias_name | string | X | √ |  | 
| last_updated_date | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| alias_arn | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| target_key_id | string | X | √ |  | 


