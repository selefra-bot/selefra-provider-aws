# Table: aws_mq_broker_configurations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| broker_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| engine_type | string | X | √ |  | 
| id | string | X | √ |  | 
| tags | json | X | √ |  | 
| aws_mq_brokers_selefra_id | string | X | X | fk to aws_mq_brokers.selefra_id | 
| description | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| authentication_strategy | string | X | √ |  | 
| created | timestamp | X | √ |  | 
| engine_version | string | X | √ |  | 
| latest_revision | json | X | √ |  | 
| name | string | X | √ |  | 
| result_metadata | json | X | √ |  | 


