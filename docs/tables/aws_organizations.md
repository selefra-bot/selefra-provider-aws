# Table: aws_organizations

## Primary Keys 

```
account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | √ | √ |  | 
| feature_set | string | X | √ |  | 
| master_account_email | string | X | √ |  | 
| master_account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| available_policy_types | json | X | √ |  | 
| id | string | X | √ |  | 
| master_account_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


