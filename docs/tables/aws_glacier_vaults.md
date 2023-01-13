# Table: aws_glacier_vaults

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| vault_arn | string | X | √ |  | 
| vault_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| last_inventory_date | string | X | √ |  | 
| creation_date | string | X | √ |  | 
| number_of_archives | big_int | X | √ |  | 
| size_in_bytes | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


