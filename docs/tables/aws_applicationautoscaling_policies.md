# Table: aws_applicationautoscaling_policies

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| policy_type | string | X | √ |  | 
| alarms | json | X | √ |  | 
| target_tracking_scaling_policy_configuration | json | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| scalable_dimension | string | X | √ |  | 
| step_scaling_policy_configuration | json | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| policy_arn | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| service_namespace | string | X | √ |  | 


