# Table: aws_lightsail_alarms

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| period | big_int | X | √ |  | 
| resource_type | string | X | √ |  | 
| statistic | string | X | √ |  | 
| account_id | string | X | √ |  | 
| comparison_operator | string | X | √ |  | 
| contact_protocols | string_array | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| evaluation_periods | big_int | X | √ |  | 
| location | json | X | √ |  | 
| name | string | X | √ |  | 
| notification_triggers | string_array | X | √ |  | 
| arn | string | X | √ |  | 
| state | string | X | √ |  | 
| treat_missing_data | string | X | √ |  | 
| unit | string | X | √ |  | 
| datapoints_to_alarm | big_int | X | √ |  | 
| metric_name | string | X | √ |  | 
| monitored_resource_info | json | X | √ |  | 
| notification_enabled | bool | X | √ |  | 
| support_code | string | X | √ |  | 
| threshold | float | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 


