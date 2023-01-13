# Table: aws_rds_db_security_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| e_c2_security_groups | json | X | √ |  | 
| vpc_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| db_security_group_arn | string | X | √ |  | 
| db_security_group_description | string | X | √ |  | 
| db_security_group_name | string | X | √ |  | 
| ip_ranges | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


