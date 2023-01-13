# Table: aws_iot_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| policies | string_array | X | √ |  | 
| previous_owned_by | string | X | √ |  | 
| validity | json | X | √ |  | 
| arn | string | √ | √ |  | 
| certificate_pem | string | X | √ |  | 
| customer_version | big_int | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| ca_certificate_id | string | X | √ |  | 
| certificate_id | string | X | √ |  | 
| generation_id | string | X | √ |  | 
| status | string | X | √ |  | 
| transfer_data | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| certificate_arn | string | X | √ |  | 
| certificate_mode | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| owned_by | string | X | √ |  | 


