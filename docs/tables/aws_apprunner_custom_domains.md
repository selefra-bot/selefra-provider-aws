# Table: aws_apprunner_custom_domains

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| enable_www_subdomain | bool | X | √ |  | 
| domain_name | string | X | √ |  | 
| status | string | X | √ |  | 
| certificate_validation_records | json | X | √ |  | 
| aws_apprunner_services_selefra_id | string | X | X | fk to aws_apprunner_services.selefra_id | 


