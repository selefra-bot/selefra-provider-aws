# Table: aws_ec2_subnets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| map_customer_owned_ip_on_launch | bool | X | √ |  | 
| private_dns_name_options_on_launch | json | X | √ |  | 
| subnet_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| availability_zone_id | string | X | √ |  | 
| cidr_block | string | X | √ |  | 
| default_for_az | bool | X | √ |  | 
| enable_lni_at_device_index | big_int | X | √ |  | 
| vpc_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| customer_owned_ipv4_pool | string | X | √ |  | 
| assign_ipv6_address_on_creation | bool | X | √ |  | 
| available_ip_address_count | big_int | X | √ |  | 
| ipv6_cidr_block_association_set | json | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| state | string | X | √ |  | 
| subnet_arn | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| enable_dns64 | bool | X | √ |  | 
| ipv6_native | bool | X | √ |  | 
| map_public_ip_on_launch | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


