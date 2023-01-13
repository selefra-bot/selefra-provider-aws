# Table: aws_iot_things

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| principals | string_array | X | √ |  | 
| thing_arn | string | X | √ |  | 
| version | big_int | X | √ |  | 
| arn | string | √ | √ |  | 
| attributes | json | X | √ |  | 
| thing_name | string | X | √ |  | 
| thing_type_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


