# Table: aws_autoscaling_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| load_balancers | json | X | √ |  | 
| load_balancer_target_groups | json | X | √ |  | 
| arn | string | √ | √ |  | 
| auto_scaling_group | json | X | √ |  | 
| notification_configurations | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


