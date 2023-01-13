# Table: aws_cloudwatchlogs_metric_filters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| creation_time | big_int | X | √ |  | 
| filter_pattern | string | X | √ |  | 
| log_group_name | string | X | √ |  | 
| metric_transformations | json | X | √ |  | 
| region | string | X | √ |  | 
| filter_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


