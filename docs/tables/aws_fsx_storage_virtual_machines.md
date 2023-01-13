# Table: aws_fsx_storage_virtual_machines

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| active_directory_configuration | json | X | √ |  | 
| file_system_id | string | X | √ |  | 
| lifecycle_transition_reason | json | X | √ |  | 
| resource_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| lifecycle | string | X | √ |  | 
| name | string | X | √ |  | 
| endpoints | json | X | √ |  | 
| root_volume_security_style | string | X | √ |  | 
| storage_virtual_machine_id | string | X | √ |  | 
| subtype | string | X | √ |  | 
| arn | string | √ | √ |  | 
| uuid | string | X | √ |  | 


