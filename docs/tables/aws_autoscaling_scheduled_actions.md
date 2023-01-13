# Table: aws_autoscaling_scheduled_actions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| auto_scaling_group_name | string | X | √ |  | 
| end_time | timestamp | X | √ |  | 
| time_zone | string | X | √ |  | 
| account_id | string | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| max_size | big_int | X | √ |  | 
| scheduled_action_arn | string | X | √ |  | 
| scheduled_action_name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| desired_capacity | big_int | X | √ |  | 
| min_size | big_int | X | √ |  | 
| recurrence | string | X | √ |  | 
| time | timestamp | X | √ |  | 
| region | string | X | √ |  | 


