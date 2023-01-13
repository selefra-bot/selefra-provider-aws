# Table: aws_ec2_transit_gateway_peering_attachments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| options | json | X | √ |  | 
| transit_gateway_attachment_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| accepter_tgw_info | json | X | √ |  | 
| region | string | X | √ |  | 
| accepter_transit_gateway_attachment_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| requester_tgw_info | json | X | √ |  | 
| state | string | X | √ |  | 
| status | json | X | √ |  | 
| account_id | string | X | √ |  | 
| transit_gateway_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| aws_ec2_transit_gateways_selefra_id | string | X | X | fk to aws_ec2_transit_gateways.selefra_id | 


