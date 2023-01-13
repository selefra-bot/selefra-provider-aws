# Table: aws_apprunner_services

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| instance_configuration | json | X | √ |  | 
| source_configuration | json | X | √ |  | 
| deleted_at | timestamp | X | √ |  | 
| health_check_configuration | json | X | √ |  | 
| region | string | X | √ |  | 
| service_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| network_configuration | json | X | √ |  | 
| encryption_configuration | json | X | √ |  | 
| service_url | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| auto_scaling_configuration_summary | json | X | √ |  | 
| service_id | string | X | √ |  | 
| service_name | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| observability_configuration | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


