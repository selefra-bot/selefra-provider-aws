# Table: aws_iam_user_groups

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| user_arn | string | X | √ |  | 
| user_id | string | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| group_name | string | X | √ |  | 
| path | string | X | √ |  | 
| aws_iam_users_selefra_id | string | X | X | fk to aws_iam_users.selefra_id | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| group_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


