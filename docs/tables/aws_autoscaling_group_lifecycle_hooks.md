# Table: aws_autoscaling_group_lifecycle_hooks

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| notification_metadata | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| default_result | string | X | √ |  | 
| lifecycle_hook_name | string | X | √ |  | 
| lifecycle_transition | string | X | √ |  | 
| account_id | string | X | √ |  | 
| group_arn | string | X | √ |  | 
| heartbeat_timeout | big_int | X | √ |  | 
| notification_target_arn | string | X | √ |  | 
| aws_autoscaling_groups_selefra_id | string | X | X | fk to aws_autoscaling_groups.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| auto_scaling_group_name | string | X | √ |  | 
| global_timeout | big_int | X | √ |  | 
| region | string | X | √ |  | 


