# Table: aws_cloudformation_stack_resources

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_type | string | X | √ |  | 
| module_info | json | X | √ |  | 
| physical_resource_id | string | X | √ |  | 
| aws_cloudformation_stacks_selefra_id | string | X | X | fk to aws_cloudformation_stacks.selefra_id | 
| region | string | X | √ |  | 
| last_updated_timestamp | timestamp | X | √ |  | 
| resource_status | string | X | √ |  | 
| drift_information | json | X | √ |  | 
| resource_status_reason | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| logical_resource_id | string | X | √ |  | 


