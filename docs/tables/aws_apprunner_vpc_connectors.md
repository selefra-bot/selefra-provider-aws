# Table: aws_apprunner_vpc_connectors

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| deleted_at | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| subnets | string_array | X | √ |  | 
| vpc_connector_revision | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| vpc_connector_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| vpc_connector_arn | string | X | √ |  | 


