# Table: aws_apprunner_vpc_ingress_connections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| source_account_id | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| service_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| vpc_ingress_connection_arn | string | X | √ |  | 
| vpc_ingress_connection_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| deleted_at | timestamp | X | √ |  | 
| domain_name | string | X | √ |  | 
| ingress_vpc_configuration | json | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 


