# Table: aws_directconnect_lags

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_logical_device_id | string | X | √ |  | 
| connections | json | X | √ |  | 
| lag_id | string | X | √ |  | 
| lag_name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| id | string | X | √ |  | 
| tags | json | X | √ |  | 
| aws_device | string | X | √ |  | 
| mac_sec_keys | json | X | √ |  | 
| provider_name | string | X | √ |  | 
| encryption_mode | string | X | √ |  | 
| jumbo_frame_capable | bool | X | √ |  | 
| minimum_links | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| allows_hosted_connections | bool | X | √ |  | 
| aws_device_v2 | string | X | √ |  | 
| connections_bandwidth | string | X | √ |  | 
| region | string | X | √ |  | 
| number_of_connections | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| owner_account | string | X | √ |  | 
| has_logical_redundancy | string | X | √ |  | 
| lag_state | string | X | √ |  | 
| location | string | X | √ |  | 
| mac_sec_capable | bool | X | √ |  | 


