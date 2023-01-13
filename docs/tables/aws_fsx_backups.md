# Table: aws_fsx_backups

## Primary Keys 

```
account_id, region, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| source_backup_region | string | X | √ |  | 
| id | string | X | √ |  | 
| directory_information | json | X | √ |  | 
| failure_details | json | X | √ |  | 
| volume | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| source_backup_id | string | X | √ |  | 
| lifecycle | string | X | √ |  | 
| type | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| resource_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| backup_id | string | X | √ |  | 
| file_system | json | X | √ |  | 
| progress_percent | big_int | X | √ |  | 
| resource_type | string | X | √ |  | 


