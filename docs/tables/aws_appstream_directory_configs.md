# Table: aws_appstream_directory_configs

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| directory_name | string | X | √ |  | 
| certificate_based_auth_properties | json | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| organizational_unit_distinguished_names | string_array | X | √ |  | 
| service_account_credentials | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


