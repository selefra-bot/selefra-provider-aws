# Table: aws_iam_role_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| policy_document | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| role_name | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| aws_iam_roles_selefra_id | string | X | X | fk to aws_iam_roles.selefra_id | 
| selefra_id | string | √ | √ | random id | 


