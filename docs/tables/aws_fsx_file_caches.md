# Table: aws_fsx_file_caches

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kms_key_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| data_repository_association_ids | string_array | X | √ |  | 
| file_cache_id | string | X | √ |  | 
| file_cache_type_version | string | X | √ |  | 
| network_interface_ids | string_array | X | √ |  | 
| storage_capacity | big_int | X | √ |  | 
| vpc_id | string | X | √ |  | 
| region | string | X | √ |  | 
| dns_name | string | X | √ |  | 
| lustre_configuration | json | X | √ |  | 
| lifecycle | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| resource_arn | string | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| failure_details | json | X | √ |  | 
| file_cache_type | string | X | √ |  | 


