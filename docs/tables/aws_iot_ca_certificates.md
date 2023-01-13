# Table: aws_iot_ca_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| certificate_pem | string | X | √ |  | 
| owned_by | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| auto_registration_status | string | X | √ |  | 
| certificate_mode | string | X | √ |  | 
| generation_id | string | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| validity | json | X | √ |  | 
| certificates | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| certificate_arn | string | X | √ |  | 
| certificate_id | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| customer_version | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


