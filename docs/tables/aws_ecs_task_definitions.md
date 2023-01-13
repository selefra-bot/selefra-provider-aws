# Table: aws_ecs_task_definitions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| ephemeral_storage | json | X | √ |  | 
| execution_role_arn | string | X | √ |  | 
| pid_mode | string | X | √ |  | 
| inference_accelerators | json | X | √ |  | 
| placement_constraints | json | X | √ |  | 
| registered_by | string | X | √ |  | 
| region | string | X | √ |  | 
| ipc_mode | string | X | √ |  | 
| proxy_configuration | json | X | √ |  | 
| revision | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| compatibilities | string_array | X | √ |  | 
| family | string | X | √ |  | 
| registered_at | timestamp | X | √ |  | 
| volumes | json | X | √ |  | 
| deregistered_at | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| task_definition_arn | string | X | √ |  | 
| container_definitions | json | X | √ |  | 
| cpu | string | X | √ |  | 
| requires_compatibilities | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| runtime_platform | json | X | √ |  | 
| memory | string | X | √ |  | 
| network_mode | string | X | √ |  | 
| requires_attributes | json | X | √ |  | 
| task_role_arn | string | X | √ |  | 


