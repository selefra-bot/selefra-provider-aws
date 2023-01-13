# Table: aws_xray_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| filter_expression | string | X | √ |  | 
| group_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| group_arn | string | X | √ |  | 
| insights_configuration | json | X | √ |  | 


