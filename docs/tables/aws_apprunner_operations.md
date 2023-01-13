# Table: aws_apprunner_operations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| started_at | timestamp | X | √ |  | 
| target_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| ended_at | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| type | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| aws_apprunner_services_selefra_id | string | X | X | fk to aws_apprunner_services.selefra_id | 


