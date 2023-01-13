# Table: aws_directconnect_virtual_interfaces

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_logical_device_id | string | X | √ |  | 
| virtual_interface_id | string | X | √ |  | 
| virtual_interface_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | X | √ |  | 
| amazon_side_asn | big_int | X | √ |  | 
| bgp_peers | json | X | √ |  | 
| route_filter_prefixes | json | X | √ |  | 
| virtual_interface_name | string | X | √ |  | 
| address_family | string | X | √ |  | 
| auth_key | string | X | √ |  | 
| connection_id | string | X | √ |  | 
| vlan | big_int | X | √ |  | 
| region | string | X | √ |  | 
| amazon_address | string | X | √ |  | 
| direct_connect_gateway_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| customer_address | string | X | √ |  | 
| site_link_enabled | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| customer_router_config | string | X | √ |  | 
| owner_account | string | X | √ |  | 
| virtual_interface_state | string | X | √ |  | 
| virtual_gateway_id | string | X | √ |  | 
| asn | big_int | X | √ |  | 
| aws_device_v2 | string | X | √ |  | 
| jumbo_frame_capable | bool | X | √ |  | 
| location | string | X | √ |  | 
| mtu | big_int | X | √ |  | 


