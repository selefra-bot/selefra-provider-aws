# Table: aws_codebuild_projects

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| artifacts | json | X | √ |  | 
| badge | json | X | √ |  | 
| queued_timeout_in_minutes | big_int | X | √ |  | 
| webhook | json | X | √ |  | 
| account_id | string | X | √ |  | 
| concurrent_build_limit | big_int | X | √ |  | 
| environment | json | X | √ |  | 
| secondary_source_versions | json | X | √ |  | 
| logs_config | json | X | √ |  | 
| project_visibility | string | X | √ |  | 
| secondary_artifacts | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| build_batch_config | json | X | √ |  | 
| file_system_locations | json | X | √ |  | 
| cache | json | X | √ |  | 
| public_project_alias | string | X | √ |  | 
| secondary_sources | json | X | √ |  | 
| timeout_in_minutes | big_int | X | √ |  | 
| arn | string | X | √ |  | 
| created | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| encryption_key | string | X | √ |  | 
| name | string | X | √ |  | 
| source | json | X | √ |  | 
| vpc_config | json | X | √ |  | 
| last_modified | timestamp | X | √ |  | 
| resource_access_role | string | X | √ |  | 
| service_role | string | X | √ |  | 
| source_version | string | X | √ |  | 
| tags | json | X | √ |  | 


