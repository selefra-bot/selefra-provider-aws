# Table: aws_directconnect_gateway_associations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| association_id | string | X | √ |  | 
| direct_connect_gateway_owner_account | string | X | √ |  | 
| virtual_gateway_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| associated_gateway | json | X | √ |  | 
| state_change_error | string | X | √ |  | 
| virtual_gateway_owner_account | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| gateway_id | string | X | √ |  | 
| allowed_prefixes_to_direct_connect_gateway | json | X | √ |  | 
| association_state | string | X | √ |  | 
| direct_connect_gateway_id | string | X | √ |  | 
| aws_directconnect_gateways_selefra_id | string | X | X | fk to aws_directconnect_gateways.selefra_id | 
| region | string | X | √ |  | 
| gateway_arn | string | X | √ |  | 
| virtual_gateway_region | string | X | √ |  | 


