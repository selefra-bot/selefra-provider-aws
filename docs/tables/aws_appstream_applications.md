# Table: aws_appstream_applications

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| display_name | string | X | √ |  | 
| icon_url | string | X | √ |  | 
| launch_parameters | string | X | √ |  | 
| name | string | X | √ |  | 
| platforms | string_array | X | √ |  | 
| working_directory | string | X | √ |  | 
| description | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| icon_s3_location | json | X | √ |  | 
| instance_families | string_array | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| metadata | json | X | √ |  | 
| app_block_arn | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| launch_path | string | X | √ |  | 


