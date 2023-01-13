# Table: aws_iot_billing_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| billing_group_arn | string | X | √ |  | 
| billing_group_id | string | X | √ |  | 
| billing_group_properties | json | X | √ |  | 
| version | big_int | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| things_in_group | string_array | X | √ |  | 
| billing_group_metadata | json | X | √ |  | 
| billing_group_name | string | X | √ |  | 


