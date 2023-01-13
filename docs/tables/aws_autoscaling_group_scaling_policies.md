# Table: aws_autoscaling_group_scaling_policies

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| adjustment_type | string | X | √ |  | 
| min_adjustment_magnitude | big_int | X | √ |  | 
| predictive_scaling_configuration | json | X | √ |  | 
| cooldown | big_int | X | √ |  | 
| estimated_instance_warmup | big_int | X | √ |  | 
| policy_arn | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| alarms | json | X | √ |  | 
| auto_scaling_group_name | string | X | √ |  | 
| policy_type | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| min_adjustment_step | big_int | X | √ |  | 
| target_tracking_configuration | json | X | √ |  | 
| aws_autoscaling_groups_selefra_id | string | X | X | fk to aws_autoscaling_groups.selefra_id | 
| step_adjustments | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| group_arn | string | X | √ |  | 
| metric_aggregation_type | string | X | √ |  | 
| scaling_adjustment | big_int | X | √ |  | 


