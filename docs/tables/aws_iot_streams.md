# Table: aws_iot_streams

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| files | json | X | √ |  | 
| stream_arn | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| stream_id | string | X | √ |  | 
| stream_version | big_int | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created_at | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| last_updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


