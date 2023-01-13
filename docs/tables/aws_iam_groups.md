# Table: aws_iam_groups

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| policies | json | X | √ |  | 
| id | string | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| group_id | string | X | √ |  | 
| group_name | string | X | √ |  | 
| path | string | X | √ |  | 
| arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


