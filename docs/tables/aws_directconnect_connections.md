# Table: aws_directconnect_connections

## Primary Keys 

```
arn, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| vlan | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| location | string | X | √ |  | 
| mac_sec_capable | bool | X | √ |  | 
| tags | json | X | √ |  | 
| has_logical_redundancy | string | X | √ |  | 
| owner_account | string | X | √ |  | 
| port_encryption_status | string | X | √ |  | 
| provider_name | string | X | √ |  | 
| connection_id | string | X | √ |  | 
| connection_name | string | X | √ |  | 
| connection_state | string | X | √ |  | 
| encryption_mode | string | X | √ |  | 
| loa_issue_time | timestamp | X | √ |  | 
| mac_sec_keys | json | X | √ |  | 
| arn | string | X | √ |  | 
| aws_device_v2 | string | X | √ |  | 
| aws_logical_device_id | string | X | √ |  | 
| jumbo_frame_capable | bool | X | √ |  | 
| lag_id | string | X | √ |  | 
| partner_name | string | X | √ |  | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| aws_device | string | X | √ |  | 
| bandwidth | string | X | √ |  | 


