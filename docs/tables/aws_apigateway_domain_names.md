# Table: aws_apigateway_domain_names

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| certificate_name | string | X | √ |  | 
| regional_certificate_arn | string | X | √ |  | 
| regional_hosted_zone_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| certificate_arn | string | X | √ |  | 
| domain_name_status | string | X | √ |  | 
| ownership_verification_certificate_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 
| certificate_upload_date | timestamp | X | √ |  | 
| distribution_domain_name | string | X | √ |  | 
| distribution_hosted_zone_id | string | X | √ |  | 
| domain_name_status_message | string | X | √ |  | 
| regional_certificate_name | string | X | √ |  | 
| regional_domain_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| domain_name | string | X | √ |  | 
| endpoint_configuration | json | X | √ |  | 
| mutual_tls_authentication | json | X | √ |  | 
| security_policy | string | X | √ |  | 


