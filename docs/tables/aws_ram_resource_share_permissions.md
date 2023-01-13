# Table: aws_ram_resource_share_permissions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_ram_resource_shares_selefra_id | string | X | X | fk to aws_ram_resource_shares.selefra_id | 
| creation_time | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| is_resource_type_default | bool | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| version | string | X | √ |  | 
| arn | string | X | √ |  | 
| default_version | bool | X | √ |  | 
| resource_type | string | X | √ |  | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| permission | json | X | √ |  | 


