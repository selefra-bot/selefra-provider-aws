# Table: aws_docdb_cluster_parameter_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| parameters | json | X | √ |  | 
| db_cluster_parameter_group_name | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| db_parameter_group_family | string | X | √ |  | 
| db_cluster_parameter_group_arn | string | X | √ |  | 
| description | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


