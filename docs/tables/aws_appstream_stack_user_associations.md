# Table: aws_appstream_stack_user_associations

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| stack_name | string | X | √ |  | 
| user_name | string | X | √ |  | 
| authentication_type | string | X | √ |  | 
| send_email_notification | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_appstream_stacks_selefra_id | string | X | X | fk to aws_appstream_stacks.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


