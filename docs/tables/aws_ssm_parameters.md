# Table: aws_ssm_parameters

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policies | json | X | √ |  | 
| type | string | X | √ |  | 
| name | string | X | √ |  | 
| allowed_pattern | string | X | √ |  | 
| key_id | string | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| description | string | X | √ |  | 
| last_modified_user | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| data_type | string | X | √ |  | 
| tier | string | X | √ |  | 
| version | big_int | X | √ |  | 


