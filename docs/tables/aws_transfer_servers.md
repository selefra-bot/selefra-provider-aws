# Table: aws_transfer_servers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| endpoint_type | string | X | √ |  | 
| host_key_fingerprint | string | X | √ |  | 
| identity_provider_details | json | X | √ |  | 
| logging_role | string | X | √ |  | 
| protocol_details | json | X | √ |  | 
| workflow_details | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| certificate | string | X | √ |  | 
| identity_provider_type | string | X | √ |  | 
| pre_authentication_login_banner | string | X | √ |  | 
| server_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| post_authentication_login_banner | string | X | √ |  | 
| user_count | big_int | X | √ |  | 
| security_policy_name | string | X | √ |  | 
| state | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| domain | string | X | √ |  | 
| endpoint_details | json | X | √ |  | 
| protocols | string_array | X | √ |  | 


