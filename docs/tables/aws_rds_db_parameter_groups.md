# Table: aws_rds_db_parameter_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| db_parameter_group_arn | string | X | √ |  | 
| db_parameter_group_family | string | X | √ |  | 
| db_parameter_group_name | string | X | √ |  | 
| description | string | X | √ |  | 


