# Table: aws_iam_policies

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| tags | json | X | √ |  | 
| update_date | timestamp | X | √ |  | 
| arn | string | X | √ |  | 
| description | string | X | √ |  | 
| is_attachable | bool | X | √ |  | 
| policy_id | string | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| default_version_id | string | X | √ |  | 
| path | string | X | √ |  | 
| permissions_boundary_usage_count | big_int | X | √ |  | 
| policy_name | string | X | √ |  | 
| policy_version_list | json | X | √ |  | 
| attachment_count | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


