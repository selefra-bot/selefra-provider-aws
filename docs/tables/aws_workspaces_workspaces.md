# Table: aws_workspaces_workspaces

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| workspace_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| computer_name | string | X | √ |  | 
| directory_id | string | X | √ |  | 
| error_code | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| user_name | string | X | √ |  | 
| bundle_id | string | X | √ |  | 
| error_message | string | X | √ |  | 
| ip_address | string | X | √ |  | 
| related_workspaces | json | X | √ |  | 
| root_volume_encryption_enabled | bool | X | √ |  | 
| volume_encryption_key | string | X | √ |  | 
| account_id | string | X | √ |  | 
| modification_states | json | X | √ |  | 
| user_volume_encryption_enabled | bool | X | √ |  | 
| state | string | X | √ |  | 
| workspace_properties | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


