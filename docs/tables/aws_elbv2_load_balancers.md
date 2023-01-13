# Table: aws_elbv2_load_balancers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| dns_name | string | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| state | json | X | √ |  | 
| web_acl_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| canonical_hosted_zone_id | string | X | √ |  | 
| load_balancer_name | string | X | √ |  | 
| scheme | string | X | √ |  | 
| type | string | X | √ |  | 
| customer_owned_ipv4_pool | string | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| load_balancer_arn | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| vpc_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| availability_zones | json | X | √ |  | 


