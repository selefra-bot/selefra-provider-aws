# Table: aws_backup_vault_recovery_points

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| composite_member_identifier | string | X | √ |  | 
| created_by | json | X | √ |  | 
| last_restore_time | timestamp | X | √ |  | 
| status_message | string | X | √ |  | 
| region | string | X | √ |  | 
| backup_vault_name | string | X | √ |  | 
| resource_arn | string | X | √ |  | 
| aws_backup_vaults_selefra_id | string | X | X | fk to aws_backup_vaults.selefra_id | 
| vault_arn | string | X | √ |  | 
| is_encrypted | bool | X | √ |  | 
| completion_date | timestamp | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| recovery_point_arn | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| status | string | X | √ |  | 
| backup_size_in_bytes | big_int | X | √ |  | 
| backup_vault_arn | string | X | √ |  | 
| lifecycle | json | X | √ |  | 
| source_backup_vault_arn | string | X | √ |  | 
| iam_role_arn | string | X | √ |  | 
| is_parent | bool | X | √ |  | 
| arn | string | √ | √ |  | 
| encryption_key_arn | string | X | √ |  | 
| calculated_lifecycle | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| parent_recovery_point_arn | string | X | √ |  | 


