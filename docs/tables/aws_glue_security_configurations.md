# Table: aws_glue_security_configurations

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| name | string | X | √ |  | 
| created_time_stamp | timestamp | X | √ |  | 
| encryption_configuration | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


