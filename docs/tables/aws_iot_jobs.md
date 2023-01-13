# Table: aws_iot_jobs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| timeout_config | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| is_concurrent | bool | X | √ |  | 
| job_process_details | json | X | √ |  | 
| namespace_id | string | X | √ |  | 
| status | string | X | √ |  | 
| target_selection | string | X | √ |  | 
| targets | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| description | string | X | √ |  | 
| document_parameters | json | X | √ |  | 
| job_executions_rollout_config | json | X | √ |  | 
| job_id | string | X | √ |  | 
| region | string | X | √ |  | 
| completed_at | timestamp | X | √ |  | 
| arn | string | √ | √ |  | 
| comment | string | X | √ |  | 
| presigned_url_config | json | X | √ |  | 
| scheduling_config | json | X | √ |  | 
| account_id | string | X | √ |  | 
| job_arn | string | X | √ |  | 
| last_updated_at | timestamp | X | √ |  | 
| job_executions_retry_config | json | X | √ |  | 
| job_template_arn | string | X | √ |  | 
| reason_code | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| force_canceled | bool | X | √ |  | 
| abort_config | json | X | √ |  | 


