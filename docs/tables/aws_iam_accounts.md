# Table: aws_iam_accounts

## Primary Keys 

```
account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| groups_quota | big_int | X | √ |  | 
| server_certificates_quota | big_int | X | √ |  | 
| group_policy_size_quota | big_int | X | √ |  | 
| aliases | string_array | X | √ |  | 
| groups | big_int | X | √ |  | 
| user_policy_size_quota | big_int | X | √ |  | 
| attached_policies_per_group_quota | big_int | X | √ |  | 
| versions_per_policy_quota | big_int | X | √ |  | 
| access_keys_per_user_quota | big_int | X | √ |  | 
| mfa_devices | big_int | X | √ |  | 
| users | big_int | X | √ |  | 
| server_certificates | big_int | X | √ |  | 
| account_signing_certificates_present | bool | X | √ |  | 
| policy_versions_in_use_quota | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| groups_per_user_quota | big_int | X | √ |  | 
| account_access_keys_present | bool | X | √ |  | 
| policy_size_quota | big_int | X | √ |  | 
| policy_versions_in_use | big_int | X | √ |  | 
| account_id | string | √ | √ |  | 
| users_quota | big_int | X | √ |  | 
| signing_certificates_per_user_quota | big_int | X | √ |  | 
| attached_policies_per_role_quota | big_int | X | √ |  | 
| global_endpoint_token_version | big_int | X | √ |  | 
| mfa_devices_in_use | big_int | X | √ |  | 
| account_mfa_enabled | bool | X | √ |  | 
| attached_policies_per_user_quota | big_int | X | √ |  | 
| policies | big_int | X | √ |  | 
| policies_quota | big_int | X | √ |  | 


