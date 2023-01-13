# Table: aws_cloudformation_stacks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enable_termination_protection | bool | X | √ |  | 
| drift_information | json | X | √ |  | 
| notification_ar_ns | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| stack_name | string | X | √ |  | 
| disable_rollback | bool | X | √ |  | 
| parameters | json | X | √ |  | 
| root_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ |  | 
| stack_status | string | X | √ |  | 
| capabilities | string_array | X | √ |  | 
| description | string | X | √ |  | 
| timeout_in_minutes | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| deletion_time | timestamp | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| parent_id | string | X | √ |  | 
| stack_id | string | X | √ |  | 
| stack_status_reason | string | X | √ |  | 
| id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| change_set_id | string | X | √ |  | 
| outputs | json | X | √ |  | 
| role_arn | string | X | √ |  | 
| rollback_configuration | json | X | √ |  | 
| region | string | X | √ |  | 


