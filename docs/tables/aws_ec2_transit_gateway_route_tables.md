# Table: aws_ec2_transit_gateway_route_tables

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| default_propagation_route_table | bool | X | √ |  | 
| aws_ec2_transit_gateways_selefra_id | string | X | X | fk to aws_ec2_transit_gateways.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| transit_gateway_arn | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| transit_gateway_id | string | X | √ |  | 
| transit_gateway_route_table_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| default_association_route_table | bool | X | √ |  | 


