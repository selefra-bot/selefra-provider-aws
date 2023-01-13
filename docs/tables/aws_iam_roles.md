# Table: aws_iam_roles

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| role_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| policies | json | X | √ |  | 
| permissions_boundary | json | X | √ |  | 
| role_name | string | X | √ |  | 
| max_session_duration | big_int | X | √ |  | 
| role_last_used | json | X | √ |  | 
| id | string | X | √ |  | 
| assume_role_policy_document | string | X | √ |  | 
| arn | string | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| path | string | X | √ |  | 
| description | string | X | √ |  | 


