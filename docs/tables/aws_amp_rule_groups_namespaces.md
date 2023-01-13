# Table: aws_amp_rule_groups_namespaces

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| status | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| workspace_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| data | int_array | X | √ |  | 
| modified_at | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| aws_amp_workspaces_selefra_id | string | X | X | fk to aws_amp_workspaces.selefra_id | 


