# Table: aws_kms_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| customer_master_key_spec | string | X | √ |  | 
| encryption_algorithms | string_array | X | √ |  | 
| region | string | X | √ |  | 
| multi_region | bool | X | √ |  | 
| tags | json | X | √ |  | 
| key_id | string | X | √ |  | 
| description | string | X | √ |  | 
| key_manager | string | X | √ |  | 
| signing_algorithms | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| aws_account_id | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| deletion_date | timestamp | X | √ |  | 
| multi_region_configuration | json | X | √ |  | 
| rotation_enabled | bool | X | √ |  | 
| enabled | bool | X | √ |  | 
| replica_keys | json | X | √ |  | 
| cloud_hsm_cluster_id | string | X | √ |  | 
| key_spec | string | X | √ |  | 
| mac_algorithms | string_array | X | √ |  | 
| xks_key_configuration | json | X | √ |  | 
| arn | string | X | √ |  | 
| expiration_model | string | X | √ |  | 
| key_state | string | X | √ |  | 
| pending_deletion_window_in_days | big_int | X | √ |  | 
| valid_to | timestamp | X | √ |  | 
| custom_key_store_id | string | X | √ |  | 
| key_usage | string | X | √ |  | 
| origin | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


