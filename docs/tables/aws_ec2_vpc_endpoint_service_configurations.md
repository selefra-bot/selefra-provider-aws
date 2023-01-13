# Table: aws_ec2_vpc_endpoint_service_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| supported_ip_address_types | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| manages_vpc_endpoints | bool | X | √ |  | 
| network_load_balancer_arns | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| base_endpoint_dns_names | string_array | X | √ |  | 
| service_type | json | X | √ |  | 
| service_name | string | X | √ |  | 
| region | string | X | √ |  | 
| acceptance_required | bool | X | √ |  | 
| payer_responsibility | string | X | √ |  | 
| private_dns_name_configuration | json | X | √ |  | 
| service_id | string | X | √ |  | 
| gateway_load_balancer_arns | string_array | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| service_state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


