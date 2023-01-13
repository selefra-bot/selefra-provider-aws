# Table: aws_ec2_vpc_endpoints

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subnet_ids | string_array | X | √ |  | 
| vpc_endpoint_type | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| dns_entries | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| private_dns_enabled | bool | X | √ |  | 
| dns_options | json | X | √ |  | 
| groups | json | X | √ |  | 
| network_interface_ids | string_array | X | √ |  | 
| policy_document | string | X | √ |  | 
| requester_managed | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| creation_timestamp | timestamp | X | √ |  | 
| service_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| state | string | X | √ |  | 
| route_table_ids | string_array | X | √ |  | 
| vpc_endpoint_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| last_error | json | X | √ |  | 


