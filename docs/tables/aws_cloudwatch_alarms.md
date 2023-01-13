# Table: aws_cloudwatch_alarms

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| alarm_actions | string_array | X | √ |  | 
| evaluate_low_sample_count_percentile | string | X | √ |  | 
| ok_actions | string_array | X | √ |  | 
| namespace | string | X | √ |  | 
| threshold | float | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| dimensions | json | X | √ |  | 
| insufficient_data_actions | string_array | X | √ |  | 
| alarm_name | string | X | √ |  | 
| evaluation_periods | int | X | √ |  | 
| extended_statistic | string | X | √ |  | 
| state_reason_data | string | X | √ |  | 
| account_id | string | X | √ |  | 
| actions_enabled | bool | X | √ |  | 
| evaluation_state | string | X | √ |  | 
| state_reason | string | X | √ |  | 
| alarm_description | string | X | √ |  | 
| comparison_operator | string | X | √ |  | 
| threshold_metric_id | string | X | √ |  | 
| treat_missing_data | string | X | √ |  | 
| metric_name | string | X | √ |  | 
| metrics | json | X | √ |  | 
| state_transitioned_timestamp | timestamp | X | √ |  | 
| state_updated_timestamp | timestamp | X | √ |  | 
| state_value | string | X | √ |  | 
| alarm_configuration_updated_timestamp | timestamp | X | √ |  | 
| datapoints_to_alarm | int | X | √ |  | 
| period | int | X | √ |  | 
| statistic | string | X | √ |  | 
| unit | string | X | √ |  | 


