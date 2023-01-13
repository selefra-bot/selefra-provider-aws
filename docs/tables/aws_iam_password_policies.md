# Table: aws_iam_password_policies

## Primary Keys 

```
account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policy_exists | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | √ | √ |  | 
| password_policy | json | X | √ |  | 


