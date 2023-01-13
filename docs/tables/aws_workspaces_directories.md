# Table: aws_workspaces_directories

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| workspace_access_properties | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| alias | string | X | √ |  | 
| customer_user_name | string | X | √ |  | 
| directory_id | string | X | √ |  | 
| dns_ip_addresses | string_array | X | √ |  | 
| saml_properties | json | X | √ |  | 
| arn | string | √ | √ |  | 
| registration_code | string | X | √ |  | 
| selfservice_permissions | json | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| workspace_creation_properties | json | X | √ |  | 
| directory_name | string | X | √ |  | 
| directory_type | string | X | √ |  | 
| tenancy | string | X | √ |  | 
| workspace_security_group_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| certificate_based_auth_properties | json | X | √ |  | 
| iam_role_id | string | X | √ |  | 
| ip_group_ids | string_array | X | √ |  | 
| state | string | X | √ |  | 


