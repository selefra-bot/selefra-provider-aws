# Table: aws_rds_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| certificate_arn | string | X | √ |  | 
| customer_override_valid_till | timestamp | X | √ |  | 
| valid_from | timestamp | X | √ |  | 
| valid_till | timestamp | X | √ |  | 
| arn | string | √ | √ |  | 
| certificate_identifier | string | X | √ |  | 
| certificate_type | string | X | √ |  | 
| customer_override | bool | X | √ |  | 
| thumbprint | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


