# Table: aws_apigateway_rest_apis

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 
| binary_media_types | string_array | X | √ |  | 
| minimum_compression_size | big_int | X | √ |  | 
| policy | string | X | √ |  | 
| description | string | X | √ |  | 
| disable_execute_api_endpoint | bool | X | √ |  | 
| endpoint_configuration | json | X | √ |  | 
| version | string | X | √ |  | 
| warnings | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| api_key_source | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 


