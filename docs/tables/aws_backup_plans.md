# Table: aws_backup_plans

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| backup_plan_id | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| last_execution_date | timestamp | X | √ |  | 
| version_id | string | X | √ |  | 
| backup_plan_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| deletion_date | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| advanced_backup_settings | json | X | √ |  | 
| backup_plan | json | X | √ |  | 
| creator_request_id | string | X | √ |  | 


