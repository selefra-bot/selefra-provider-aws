# Table: aws_mq_broker_configuration_revisions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| data | string | X | √ |  | 
| description | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| broker_configuration_arn | string | X | √ |  | 
| configuration_id | string | X | √ |  | 
| created | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_mq_broker_configurations_selefra_id | string | X | X | fk to aws_mq_broker_configurations.selefra_id | 
| region | string | X | √ |  | 


