# Table: aws_fsx_file_systems

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| administrative_actions | json | X | √ |  | 
| storage_type | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| failure_details | json | X | √ |  | 
| file_system_type | string | X | √ |  | 
| lustre_configuration | json | X | √ |  | 
| network_interface_ids | string_array | X | √ |  | 
| storage_capacity | big_int | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| windows_configuration | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| file_system_id | string | X | √ |  | 
| file_system_type_version | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| resource_arn | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| dns_name | string | X | √ |  | 
| lifecycle | string | X | √ |  | 
| ontap_configuration | json | X | √ |  | 
| open_zfs_configuration | json | X | √ |  | 
| vpc_id | string | X | √ |  | 


