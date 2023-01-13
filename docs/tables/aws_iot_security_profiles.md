# Table: aws_iot_security_profiles

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| behaviors | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| targets | string_array | X | √ |  | 
| additional_metrics_to_retain_v2 | json | X | √ |  | 
| arn | string | √ | √ |  | 
| alert_targets | json | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| security_profile_description | string | X | √ |  | 
| version | int | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| additional_metrics_to_retain | string_array | X | √ |  | 
| security_profile_name | string | X | √ |  | 
| result_metadata | json | X | √ |  | 


