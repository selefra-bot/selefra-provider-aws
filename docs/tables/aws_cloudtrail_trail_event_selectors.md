# Table: aws_cloudtrail_trail_event_selectors

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| aws_cloudtrail_trails_selefra_id | string | X | X | fk to aws_cloudtrail_trails.selefra_id | 
| exclude_management_event_sources | string_array | X | √ |  | 
| region | string | X | √ |  | 
| trail_arn | string | X | √ |  | 
| data_resources | json | X | √ |  | 
| include_management_events | bool | X | √ |  | 
| read_write_type | string | X | √ |  | 
| account_id | string | X | √ |  | 


