# Table: aws_lambda_function_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| revision_id | string | X | √ |  | 
| role | string | X | √ |  | 
| tracing_config | json | X | √ |  | 
| code_sha256 | string | X | √ |  | 
| description | string | X | √ |  | 
| kms_key_arn | string | X | √ |  | 
| memory_size | big_int | X | √ |  | 
| state | string | X | √ |  | 
| vpc_config | json | X | √ |  | 
| account_id | string | X | √ |  | 
| last_update_status_reason | string | X | √ |  | 
| last_update_status_reason_code | string | X | √ |  | 
| timeout | big_int | X | √ |  | 
| architectures | string_array | X | √ |  | 
| code_size | big_int | X | √ |  | 
| state_reason_code | string | X | √ |  | 
| last_modified | string | X | √ |  | 
| runtime | string | X | √ |  | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| function_arn | string | X | √ |  | 
| ephemeral_storage | json | X | √ |  | 
| image_config_response | json | X | √ |  | 
| state_reason | string | X | √ |  | 
| function_name | string | X | √ |  | 
| handler | string | X | √ |  | 
| last_update_status | string | X | √ |  | 
| master_arn | string | X | √ |  | 
| signing_profile_version_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| dead_letter_config | json | X | √ |  | 
| environment | json | X | √ |  | 
| signing_job_arn | string | X | √ |  | 
| snap_start | json | X | √ |  | 
| version | string | X | √ |  | 
| file_system_configs | json | X | √ |  | 
| layers | json | X | √ |  | 
| package_type | string | X | √ |  | 


