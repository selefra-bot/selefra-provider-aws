# Table: aws_ecs_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| registered_container_instances_count | big_int | X | √ |  | 
| status | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| attachments | json | X | √ |  | 
| configuration | json | X | √ |  | 
| default_capacity_provider_strategy | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| active_services_count | big_int | X | √ |  | 
| running_tasks_count | big_int | X | √ |  | 
| service_connect_defaults | json | X | √ |  | 
| account_id | string | X | √ |  | 
| attachments_status | string | X | √ |  | 
| capacity_providers | string_array | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| cluster_name | string | X | √ |  | 
| pending_tasks_count | big_int | X | √ |  | 
| settings | json | X | √ |  | 
| statistics | json | X | √ |  | 


