# Table: aws_glue_jobs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_on | timestamp | X | √ |  | 
| command | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| code_gen_configuration_nodes | json | X | √ |  | 
| execution_class | string | X | √ |  | 
| max_retries | big_int | X | √ |  | 
| non_overridable_arguments | json | X | √ |  | 
| security_configuration | string | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| description | string | X | √ |  | 
| glue_version | string | X | √ |  | 
| source_control_details | json | X | √ |  | 
| connections | json | X | √ |  | 
| max_capacity | float | X | √ |  | 
| timeout | big_int | X | √ |  | 
| region | string | X | √ |  | 
| notification_property | json | X | √ |  | 
| worker_type | string | X | √ |  | 
| role | string | X | √ |  | 
| allocated_capacity | big_int | X | √ |  | 
| default_arguments | json | X | √ |  | 
| execution_property | json | X | √ |  | 
| log_uri | string | X | √ |  | 
| number_of_workers | big_int | X | √ |  | 
| arn | string | √ | √ |  | 
| last_modified_on | timestamp | X | √ |  | 
| name | string | X | √ |  | 


