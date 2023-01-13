# Table: aws_secretsmanager_secrets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_accessed_date | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| deleted_date | timestamp | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| rotation_lambda_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| description | string | X | √ |  | 
| last_rotated_date | timestamp | X | √ |  | 
| primary_region | string | X | √ |  | 
| version_ids_to_stages | json | X | √ |  | 
| last_changed_date | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| next_rotation_date | timestamp | X | √ |  | 
| replication_status | json | X | √ |  | 
| rotation_enabled | bool | X | √ |  | 
| rotation_rules | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| policy | json | X | √ | `A JSON-formatted string that describes the permissions that are associated with the attached secret.` | 
| tags | json | X | √ |  | 
| owning_service | string | X | √ |  | 


