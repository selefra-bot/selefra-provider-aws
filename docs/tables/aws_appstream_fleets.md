# Table: aws_appstream_fleets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| fleet_errors | json | X | √ |  | 
| fleet_type | string | X | √ |  | 
| compute_capacity_status | json | X | √ |  | 
| name | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| idle_disconnect_timeout_in_seconds | big_int | X | √ |  | 
| region | string | X | √ |  | 
| enable_default_internet_access | bool | X | √ |  | 
| max_user_duration_in_seconds | big_int | X | √ |  | 
| stream_view | string | X | √ |  | 
| vpc_config | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| image_name | string | X | √ |  | 
| description | string | X | √ |  | 
| disconnect_timeout_in_seconds | big_int | X | √ |  | 
| iam_role_arn | string | X | √ |  | 
| platform | string | X | √ |  | 
| session_script_s3_location | json | X | √ |  | 
| instance_type | string | X | √ |  | 
| state | string | X | √ |  | 
| image_arn | string | X | √ |  | 
| usb_device_filter_strings | string_array | X | √ |  | 
| display_name | string | X | √ |  | 
| domain_join_info | json | X | √ |  | 
| max_concurrent_sessions | big_int | X | √ |  | 


