# Table: aws_ram_principals

## Primary Keys 

```
account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| external | bool | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| resource_share_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | √ | √ |  | 
| region | string | X | √ |  | 


