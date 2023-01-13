# Table: aws_athena_work_group_query_executions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| engine_version | json | X | √ |  | 
| region | string | X | √ |  | 
| execution_parameters | string_array | X | √ |  | 
| query_execution_context | json | X | √ |  | 
| query_execution_id | string | X | √ |  | 
| result_configuration | json | X | √ |  | 
| statement_type | string | X | √ |  | 
| work_group | string | X | √ |  | 
| statistics | json | X | √ |  | 
| status | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| work_group_arn | string | X | √ |  | 
| query | string | X | √ |  | 
| result_reuse_configuration | json | X | √ |  | 
| aws_athena_work_groups_selefra_id | string | X | X | fk to aws_athena_work_groups.selefra_id | 


