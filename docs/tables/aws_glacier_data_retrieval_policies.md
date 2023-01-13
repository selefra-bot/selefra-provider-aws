# Table: aws_glacier_data_retrieval_policies

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| rules | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


