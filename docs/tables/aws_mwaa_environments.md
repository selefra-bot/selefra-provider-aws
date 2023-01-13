# Table: aws_mwaa_environments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| source_bucket_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| webserver_access_mode | string | X | √ |  | 
| arn | string | X | √ |  | 
| airflow_version | string | X | √ |  | 
| min_workers | big_int | X | √ |  | 
| schedulers | big_int | X | √ |  | 
| environment_class | string | X | √ |  | 
| network_configuration | json | X | √ |  | 
| logging_configuration | json | X | √ |  | 
| plugins_s3_path | string | X | √ |  | 
| requirements_s3_path | string | X | √ |  | 
| max_workers | big_int | X | √ |  | 
| kms_key | string | X | √ |  | 
| service_role_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| last_update | json | X | √ |  | 
| plugins_s3_object_version | string | X | √ |  | 
| weekly_maintenance_window_start | string | X | √ |  | 
| status | string | X | √ |  | 
| webserver_url | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| airflow_configuration_options | json | X | √ |  | 
| execution_role_arn | string | X | √ |  | 
| name | string | X | √ |  | 
| requirements_s3_object_version | string | X | √ |  | 
| region | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| dag_s3_path | string | X | √ |  | 


