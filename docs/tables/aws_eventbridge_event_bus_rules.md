# Table: aws_eventbridge_event_bus_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| managed_by | string | X | √ |  | 
| name | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| description | string | X | √ |  | 
| event_bus_name | string | X | √ |  | 
| event_bus_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_eventbridge_event_buses_selefra_id | string | X | X | fk to aws_eventbridge_event_buses.selefra_id | 
| event_pattern | string | X | √ |  | 
| schedule_expression | string | X | √ |  | 
| state | string | X | √ |  | 


