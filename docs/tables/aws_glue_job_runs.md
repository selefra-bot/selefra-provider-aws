# Table: aws_glue_job_runs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| trigger_name | string | X | √ |  | 
| max_capacity | float | X | √ |  | 
| completed_on | timestamp | X | √ |  | 
| glue_version | string | X | √ |  | 
| last_modified_on | timestamp | X | √ |  | 
| predecessor_runs | json | X | √ |  | 
| aws_glue_jobs_selefra_id | string | X | X | fk to aws_glue_jobs.selefra_id | 
| account_id | string | X | √ |  | 
| job_run_state | string | X | √ |  | 
| log_group_name | string | X | √ |  | 
| number_of_workers | big_int | X | √ |  | 
| security_configuration | string | X | √ |  | 
| job_name | string | X | √ |  | 
| error_message | string | X | √ |  | 
| timeout | big_int | X | √ |  | 
| worker_type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| allocated_capacity | big_int | X | √ |  | 
| attempt | big_int | X | √ |  | 
| id | string | X | √ |  | 
| started_on | timestamp | X | √ |  | 
| job_arn | string | X | √ |  | 
| arguments | json | X | √ |  | 
| dpu_seconds | float | X | √ |  | 
| execution_class | string | X | √ |  | 
| execution_time | big_int | X | √ |  | 
| previous_run_id | string | X | √ |  | 
| region | string | X | √ |  | 
| notification_property | json | X | √ |  | 


