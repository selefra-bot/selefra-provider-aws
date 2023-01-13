# Table: aws_glue_ml_transform_task_runs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| execution_time | big_int | X | √ |  | 
| last_modified_on | timestamp | X | √ |  | 
| started_on | timestamp | X | √ |  | 
| task_run_id | string | X | √ |  | 
| aws_glue_ml_transforms_selefra_id | string | X | X | fk to aws_glue_ml_transforms.selefra_id | 
| completed_on | timestamp | X | √ |  | 
| properties | json | X | √ |  | 
| region | string | X | √ |  | 
| ml_transform_arn | string | X | √ |  | 
| error_string | string | X | √ |  | 
| status | string | X | √ |  | 
| transform_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| log_group_name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


