# Table: aws_athena_work_group_named_queries

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| database | string | X | √ |  | 
| query_string | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| work_group_arn | string | X | √ |  | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| named_query_id | string | X | √ |  | 
| work_group | string | X | √ |  | 
| aws_athena_work_groups_selefra_id | string | X | X | fk to aws_athena_work_groups.selefra_id | 


