# Table: aws_glue_connections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| connection_type | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| last_updated_by | string | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| connection_properties | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| physical_connection_requirements | json | X | √ |  | 
| region | string | X | √ |  | 
| description | string | X | √ |  | 
| match_criteria | string_array | X | √ |  | 


