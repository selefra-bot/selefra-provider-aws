# Table: aws_route53_hosted_zone_resource_record_sets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| multi_value_answer | bool | X | √ |  | 
| region | string | X | √ |  | 
| hosted_zone_arn | string | X | √ |  | 
| name | string | X | √ |  | 
| geo_location | json | X | √ |  | 
| alias_target | json | X | √ |  | 
| set_identifier | string | X | √ |  | 
| health_check_id | string | X | √ |  | 
| resource_records | json | X | √ |  | 
| ttl | big_int | X | √ |  | 
| failover | string | X | √ |  | 
| traffic_policy_instance_id | string | X | √ |  | 
| weight | big_int | X | √ |  | 
| aws_route53_hosted_zones_selefra_id | string | X | X | fk to aws_route53_hosted_zones.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| type | string | X | √ |  | 
| cidr_routing_config | json | X | √ |  | 


