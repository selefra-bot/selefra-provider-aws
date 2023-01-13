# Table: aws_wafv2_managed_rule_groups

## Primary Keys 

```
account_id, region, scope
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| scope | string | X | √ |  | 
| properties | json | X | √ |  | 
| versioning_supported | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| description | string | X | √ |  | 
| name | string | X | √ |  | 
| vendor_name | string | X | √ |  | 


