# Table: aws_ecr_registries

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| registry_id | string | X | √ |  | 
| replication_configuration | json | X | √ |  | 
| result_metadata | json | X | √ |  | 


