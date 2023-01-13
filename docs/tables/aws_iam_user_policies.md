# Table: aws_iam_user_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policy_document | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| user_id | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| user_name | string | X | √ |  | 
| aws_iam_users_selefra_id | string | X | X | fk to aws_iam_users.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| user_arn | string | X | √ |  | 


