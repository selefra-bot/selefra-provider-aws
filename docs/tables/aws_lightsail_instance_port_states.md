# Table: aws_lightsail_instance_port_states

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ipv6_cidrs | string_array | X | √ |  | 
| protocol | string | X | √ |  | 
| to_port | big_int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| instance_arn | string | X | √ |  | 
| from_port | big_int | X | √ |  | 
| state | string | X | √ |  | 
| aws_lightsail_instances_selefra_id | string | X | X | fk to aws_lightsail_instances.selefra_id | 
| region | string | X | √ |  | 
| cidr_list_aliases | string_array | X | √ |  | 
| cidrs | string_array | X | √ |  | 


