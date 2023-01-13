# Table: aws_wafv2_web_acls

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| web_acl | json | X | √ |  | 
| logging_configuration | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| resources_for_web_acl | string_array | X | √ |  | 
| arn | string | √ | √ |  | 


