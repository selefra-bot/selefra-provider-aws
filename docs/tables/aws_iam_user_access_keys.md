# Table: aws_iam_user_access_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| user_arn | string | X | √ |  | 
| user_id | string | X | √ |  | 
| last_used_service_name | string | X | √ |  | 
| access_key_metadata | json | X | √ |  | 
| last_rotated | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| last_used | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_iam_users_selefra_id | string | X | X | fk to aws_iam_users.selefra_id | 


