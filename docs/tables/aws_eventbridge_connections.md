# Table: aws_eventbridge_connections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| connection_state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| authorization_type | string | X | √ |  | 
| connection_arn | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| last_authorized_time | timestamp | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| name | string | X | √ |  | 
| state_reason | string | X | √ |  | 


