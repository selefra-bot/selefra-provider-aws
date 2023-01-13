# Table: aws_apigateway_client_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created_date | timestamp | X | √ |  | 
| pem_encoded_certificate | string | X | √ |  | 
| account_id | string | X | √ |  | 
| client_certificate_id | string | X | √ |  | 
| expiration_date | timestamp | X | √ |  | 


