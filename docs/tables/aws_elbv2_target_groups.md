# Table: aws_elbv2_target_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| health_check_interval_seconds | big_int | X | √ |  | 
| target_group_arn | string | X | √ |  | 
| target_group_name | string | X | √ |  | 
| target_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| health_check_enabled | bool | X | √ |  | 
| health_check_path | string | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| matcher | json | X | √ |  | 
| port | big_int | X | √ |  | 
| protocol | string | X | √ |  | 
| unhealthy_threshold_count | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| health_check_port | string | X | √ |  | 
| health_check_protocol | string | X | √ |  | 
| health_check_timeout_seconds | big_int | X | √ |  | 
| region | string | X | √ |  | 
| healthy_threshold_count | big_int | X | √ |  | 
| load_balancer_arns | string_array | X | √ |  | 
| protocol_version | string | X | √ |  | 
| vpc_id | string | X | √ |  | 


