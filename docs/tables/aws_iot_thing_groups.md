# Table: aws_iot_thing_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| thing_group_metadata | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| region | string | X | √ |  | 
| policies | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| index_name | string | X | √ |  | 
| query_string | string | X | √ |  | 
| thing_group_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| query_version | string | X | √ |  | 
| status | string | X | √ |  | 
| thing_group_arn | string | X | √ |  | 
| thing_group_name | string | X | √ |  | 
| thing_group_properties | json | X | √ |  | 
| things_in_group | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| version | big_int | X | √ |  | 


