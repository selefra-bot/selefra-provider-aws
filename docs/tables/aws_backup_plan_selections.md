# Table: aws_backup_plan_selections

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| backup_plan_id | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| selection_id | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| plan_arn | string | X | √ |  | 
| backup_selection | json | X | √ |  | 
| creator_request_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_backup_plans_selefra_id | string | X | X | fk to aws_backup_plans.selefra_id | 


