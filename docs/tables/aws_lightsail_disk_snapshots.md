# Table: aws_lightsail_disk_snapshots

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| from_disk_arn | string | X | √ |  | 
| from_instance_arn | string | X | √ |  | 
| state | string | X | √ |  | 
| progress | string | X | √ |  | 
| size_in_gb | big_int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| disk_arn | string | X | √ |  | 
| from_disk_name | string | X | √ |  | 
| is_from_auto_snapshot | bool | X | √ |  | 
| location | json | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| support_code | string | X | √ |  | 
| arn | string | X | √ |  | 
| from_instance_name | string | X | √ |  | 
| name | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| aws_lightsail_disks_selefra_id | string | X | X | fk to aws_lightsail_disks.selefra_id | 


