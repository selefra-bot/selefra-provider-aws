# Table: aws_eventbridge_api_destinations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| api_destination_state | string | X | √ |  | 
| connection_arn | string | X | √ |  | 
| http_method | string | X | √ |  | 
| invocation_endpoint | string | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| api_destination_arn | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| invocation_rate_limit_per_second | big_int | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


