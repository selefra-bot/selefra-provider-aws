# Table: aws_iam_users

## Primary Keys 

```
id, account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | X | √ |  | 
| tags | json | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| path | string | X | √ |  | 
| user_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| user_name | string | X | √ |  | 
| password_last_used | timestamp | X | √ |  | 
| permissions_boundary | json | X | √ |  | 


