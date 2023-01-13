# Table: aws_ec2_hosts

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| available_capacity | json | X | √ |  | 
| client_token | string | X | √ |  | 
| host_reservation_id | string | X | √ |  | 
| region | string | X | √ |  | 
| auto_placement | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| availability_zone_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| allocation_time | timestamp | X | √ |  | 
| allows_multiple_instance_types | string | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| state | string | X | √ |  | 
| host_id | string | X | √ |  | 
| host_properties | json | X | √ |  | 
| instances | json | X | √ |  | 
| member_of_service_linked_resource_group | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| host_recovery | string | X | √ |  | 
| release_time | timestamp | X | √ |  | 


