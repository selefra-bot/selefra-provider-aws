# Table: aws_kms_key_grants

## Primary Keys 

```
key_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| key_arn | string | √ | √ |  | 
| issuing_account | string | X | √ |  | 
| key_id | string | X | √ |  | 
| grantee_principal | string | X | √ |  | 
| name | string | X | √ |  | 
| retiring_principal | string | X | √ |  | 
| account_id | string | X | √ |  | 
| operations | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| grant_id | string | X | √ |  | 
| constraints | json | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| aws_kms_keys_selefra_id | string | X | X | fk to aws_kms_keys.selefra_id | 


