# Table: aws_stepfunctions_state_machines

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state_machine_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| tracing_configuration | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| definition | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| logging_configuration | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 
| type | string | X | √ |  | 
| label | string | X | √ |  | 


