# Table: aws_lightsail_database_events

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| message | string | X | √ |  | 
| aws_lightsail_databases_selefra_id | string | X | X | fk to aws_lightsail_databases.selefra_id | 
| region | string | X | √ |  | 
| database_arn | string | X | √ |  | 
| event_categories | string_array | X | √ |  | 
| resource | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


