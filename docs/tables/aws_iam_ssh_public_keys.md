# Table: aws_iam_ssh_public_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| user_arn | string | X | √ |  | 
| user_id | string | X | √ |  | 
| user_name | string | X | √ |  | 
| aws_iam_users_selefra_id | string | X | X | fk to aws_iam_users.selefra_id | 
| ssh_public_key_id | string | X | √ |  | 
| status | string | X | √ |  | 
| upload_date | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


