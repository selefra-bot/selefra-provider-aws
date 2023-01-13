# Table: aws_apigateway_api_keys

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| customer_id | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| stage_keys | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| last_updated_date | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| value | string | X | √ |  | 
| arn | string | √ | √ |  | 
| description | string | X | √ |  | 
| name | string | X | √ |  | 


