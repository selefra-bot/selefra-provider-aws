# Table: aws_apprunner_observability_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| observability_configuration_arn | string | X | √ |  | 
| observability_configuration_name | string | X | √ |  | 
| trace_configuration | json | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created_at | timestamp | X | √ |  | 
| deleted_at | timestamp | X | √ |  | 
| latest | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| observability_configuration_revision | big_int | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


