# Table: aws_ssoadmin_permission_sets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| inline_policy | json | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| permission_set_arn | string | X | √ |  | 
| relay_state | string | X | √ |  | 
| session_duration | string | X | √ |  | 
| aws_ssoadmin_instances_selefra_id | string | X | X | fk to aws_ssoadmin_instances.selefra_id | 
| name | string | X | √ |  | 


