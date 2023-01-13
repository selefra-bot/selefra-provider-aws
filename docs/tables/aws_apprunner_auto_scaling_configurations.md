# Table: aws_apprunner_auto_scaling_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| latest | bool | X | √ |  | 
| max_concurrency | big_int | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| auto_scaling_configuration_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created_at | timestamp | X | √ |  | 
| deleted_at | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| auto_scaling_configuration_name | string | X | √ |  | 
| auto_scaling_configuration_revision | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| max_size | big_int | X | √ |  | 
| min_size | big_int | X | √ |  | 
| account_id | string | X | √ |  | 


