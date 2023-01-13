# Table: aws_route53_hosted_zone_traffic_policy_instances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| traffic_policy_type | string | X | √ |  | 
| traffic_policy_version | big_int | X | √ |  | 
| arn | string | X | √ | `Amazon Resource Name (ARN) of the route53 hosted zone traffic policy instance.` | 
| hosted_zone_id | string | X | √ |  | 
| name | string | X | √ |  | 
| state | string | X | √ |  | 
| hosted_zone_arn | string | X | √ |  | 
| message | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_route53_hosted_zones_selefra_id | string | X | X | fk to aws_route53_hosted_zones.selefra_id | 
| ttl | big_int | X | √ |  | 
| traffic_policy_id | string | X | √ |  | 


