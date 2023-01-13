# Table: aws_wafregional_web_acls

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ | `Web ACL tags.` | 
| resources_for_web_acl | string_array | X | √ |  | 
| rules | json | X | √ |  | 
| web_acl_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| web_acl_id | string | X | √ |  | 
| metric_name | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| default_action | json | X | √ |  | 


