# Table: aws_cognito_identity_pools

## Primary Keys 

```
account_id, region, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| allow_unauthenticated_identities | bool | X | √ |  | 
| identity_pool_name | string | X | √ |  | 
| open_id_connect_provider_ar_ns | string_array | X | √ |  | 
| arn | string | X | √ |  | 
| identity_pool_id | string | X | √ |  | 
| allow_classic_flow | bool | X | √ |  | 
| cognito_identity_providers | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| developer_provider_name | string | X | √ |  | 
| supported_login_providers | json | X | √ |  | 
| saml_provider_ar_ns | string_array | X | √ |  | 
| identity_pool_tags | json | X | √ |  | 
| result_metadata | json | X | √ |  | 


