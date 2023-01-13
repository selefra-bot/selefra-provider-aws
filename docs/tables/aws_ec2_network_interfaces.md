# Table: aws_ec2_network_interfaces

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| ipv6_native | bool | X | √ |  | 
| requester_id | string | X | √ |  | 
| ipv6_address | string | X | √ |  | 
| account_id | string | X | √ |  | 
| association | json | X | √ |  | 
| interface_type | string | X | √ |  | 
| private_ip_address | string | X | √ |  | 
| ipv6_addresses | json | X | √ |  | 
| ipv6_prefixes | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| requester_managed | bool | X | √ |  | 
| source_dest_check | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| ipv4_prefixes | json | X | √ |  | 
| mac_address | string | X | √ |  | 
| private_ip_addresses | json | X | √ |  | 
| tags | json | X | √ |  | 
| network_interface_id | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| region | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| status | string | X | √ |  | 
| tag_set | json | X | √ |  | 
| availability_zone | string | X | √ |  | 
| groups | json | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| attachment | json | X | √ |  | 
| deny_all_igw_traffic | bool | X | √ |  | 
| description | string | X | √ |  | 


