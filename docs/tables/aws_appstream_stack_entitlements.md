# Table: aws_appstream_stack_entitlements

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| stack_name | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| aws_appstream_stacks_selefra_id | string | X | X | fk to aws_appstream_stacks.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| name | string | X | √ |  | 
| app_visibility | string | X | √ |  | 
| attributes | json | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 


