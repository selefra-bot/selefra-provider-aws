# Table: aws_ec2_vpc_endpoint_services

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| manages_vpc_endpoints | bool | X | √ |  | 
| owner | string | X | √ |  | 
| private_dns_names | json | X | √ |  | 
| service_type | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| base_endpoint_dns_names | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| service_id | string | X | √ |  | 
| vpc_endpoint_policy_supported | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| acceptance_required | bool | X | √ |  | 
| service_name | string | X | √ |  | 
| supported_ip_address_types | string_array | X | √ |  | 
| payer_responsibility | string | X | √ |  | 
| private_dns_name_verification_state | string | X | √ |  | 


