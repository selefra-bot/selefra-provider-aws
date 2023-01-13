# Table: aws_ec2_transit_gateway_multicast_domains

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| transit_gateway_multicast_domain_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| transit_gateway_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| aws_ec2_transit_gateways_selefra_id | string | X | X | fk to aws_ec2_transit_gateways.selefra_id | 
| options | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| state | string | X | √ |  | 
| transit_gateway_id | string | X | √ |  | 
| transit_gateway_multicast_domain_arn | string | X | √ |  | 


