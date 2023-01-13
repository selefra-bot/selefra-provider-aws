# Table: aws_ssm_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_latest_version | bool | X | √ |  | 
| last_successful_association_execution_date | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| platform_type | string | X | √ |  | 
| agent_version | string | X | √ |  | 
| association_status | string | X | √ |  | 
| computer_name | string | X | √ |  | 
| ip_address | string | X | √ |  | 
| source_id | string | X | √ |  | 
| last_ping_date_time | timestamp | X | √ |  | 
| platform_version | string | X | √ |  | 
| registration_date | timestamp | X | √ |  | 
| source_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| iam_role | string | X | √ |  | 
| last_association_execution_date | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| activation_id | string | X | √ |  | 
| association_overview | json | X | √ |  | 
| platform_name | string | X | √ |  | 
| instance_id | string | X | √ |  | 
| ping_status | string | X | √ |  | 
| resource_type | string | X | √ |  | 


