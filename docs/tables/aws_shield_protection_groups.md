# Table: aws_shield_protection_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| aggregation | string | X | √ |  | 
| pattern | string | X | √ |  | 
| protection_group_id | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| members | string_array | X | √ |  | 
| protection_group_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


