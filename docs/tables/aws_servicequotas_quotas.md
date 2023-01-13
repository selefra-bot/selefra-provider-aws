# Table: aws_servicequotas_quotas

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| global_quota | bool | X | √ |  | 
| period | json | X | √ |  | 
| quota_name | string | X | √ |  | 
| service_name | string | X | √ |  | 
| region | string | X | √ |  | 
| error_reason | json | X | √ |  | 
| quota_arn | string | X | √ |  | 
| service_code | string | X | √ |  | 
| arn | string | √ | √ |  | 
| adjustable | bool | X | √ |  | 
| quota_code | string | X | √ |  | 
| value | float | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| unit | string | X | √ |  | 
| usage_metric | json | X | √ |  | 
| aws_servicequotas_services_selefra_id | string | X | X | fk to aws_servicequotas_services.selefra_id | 


