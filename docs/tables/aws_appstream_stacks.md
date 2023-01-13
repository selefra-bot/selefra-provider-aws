# Table: aws_appstream_stacks

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| display_name | string | X | √ |  | 
| embed_host_domains | string_array | X | √ |  | 
| region | string | X | √ |  | 
| access_endpoints | json | X | √ |  | 
| description | string | X | √ |  | 
| streaming_experience_settings | json | X | √ |  | 
| feedback_url | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| application_settings | json | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| redirect_url | string | X | √ |  | 
| stack_errors | json | X | √ |  | 
| storage_connectors | json | X | √ |  | 
| user_settings | json | X | √ |  | 


