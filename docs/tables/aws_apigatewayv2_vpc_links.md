# Table: aws_apigatewayv2_vpc_links

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| vpc_link_status_message | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| vpc_link_id | string | X | √ |  | 
| vpc_link_status | string | X | √ |  | 
| vpc_link_version | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| security_group_ids | string_array | X | √ |  | 


