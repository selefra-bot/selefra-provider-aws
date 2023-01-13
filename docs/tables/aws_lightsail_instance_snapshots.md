# Table: aws_lightsail_instance_snapshots

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| from_instance_name | string | X | √ |  | 
| size_in_gb | big_int | X | √ |  | 
| from_blueprint_id | string | X | √ |  | 
| from_bundle_id | string | X | √ |  | 
| name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| is_from_auto_snapshot | bool | X | √ |  | 
| progress | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| from_instance_arn | string | X | √ |  | 
| from_attached_disks | json | X | √ |  | 
| location | json | X | √ |  | 
| resource_type | string | X | √ |  | 
| state | string | X | √ |  | 
| support_code | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| arn | string | X | √ |  | 
| tags | json | X | √ |  | 


