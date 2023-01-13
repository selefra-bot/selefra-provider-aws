# Table: aws_mq_brokers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| result_metadata | json | X | √ |  | 
| broker_id | string | X | √ |  | 
| created | timestamp | X | √ |  | 
| encryption_options | json | X | √ |  | 
| maintenance_window_start_time | json | X | √ |  | 
| pending_ldap_server_metadata | json | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| actions_required | json | X | √ |  | 
| deployment_mode | string | X | √ |  | 
| logs | json | X | √ |  | 
| pending_authentication_strategy | string | X | √ |  | 
| storage_type | string | X | √ |  | 
| tags | json | X | √ |  | 
| broker_instances | json | X | √ |  | 
| broker_name | string | X | √ |  | 
| configurations | json | X | √ |  | 
| authentication_strategy | string | X | √ |  | 
| pending_host_instance_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| broker_arn | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| pending_engine_version | string | X | √ |  | 
| broker_state | string | X | √ |  | 
| host_instance_type | string | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| engine_type | string | X | √ |  | 
| users | json | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| ldap_server_metadata | json | X | √ |  | 
| pending_security_groups | string_array | X | √ |  | 
| security_groups | string_array | X | √ |  | 


