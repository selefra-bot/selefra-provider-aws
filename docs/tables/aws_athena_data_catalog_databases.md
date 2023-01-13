# Table: aws_athena_data_catalog_databases

## Primary Keys 

```
data_catalog_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| parameters | json | X | √ |  | 
| aws_athena_data_catalogs_selefra_id | string | X | X | fk to aws_athena_data_catalogs.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| data_catalog_arn | string | √ | √ |  | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 


