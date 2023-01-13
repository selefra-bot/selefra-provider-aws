# Table: aws_ecs_cluster_tasks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| platform_family | string | X | √ |  | 
| pull_started_at | timestamp | X | √ |  | 
| started_by | string | X | √ |  | 
| stop_code | string | X | √ |  | 
| aws_ecs_clusters_selefra_id | string | X | X | fk to aws_ecs_clusters.selefra_id | 
| attributes | json | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| cpu | string | X | √ |  | 
| health_status | string | X | √ |  | 
| task_protection | json | X | √ |  | 
| connectivity_at | timestamp | X | √ |  | 
| ephemeral_storage | json | X | √ |  | 
| stopped_reason | string | X | √ |  | 
| region | string | X | √ |  | 
| desired_status | string | X | √ |  | 
| memory | string | X | √ |  | 
| platform_version | string | X | √ |  | 
| pull_stopped_at | timestamp | X | √ |  | 
| version | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| capacity_provider_name | string | X | √ |  | 
| container_instance_arn | string | X | √ |  | 
| enable_execute_command | bool | X | √ |  | 
| last_status | string | X | √ |  | 
| stopping_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| execution_stopped_at | timestamp | X | √ |  | 
| started_at | timestamp | X | √ |  | 
| arn | string | √ | √ |  | 
| inference_accelerators | json | X | √ |  | 
| task_definition_arn | string | X | √ |  | 
| containers | json | X | √ |  | 
| connectivity | string | X | √ |  | 
| group | string | X | √ |  | 
| launch_type | string | X | √ |  | 
| overrides | json | X | √ |  | 
| stopped_at | timestamp | X | √ |  | 
| task_arn | string | X | √ |  | 
| attachments | json | X | √ |  | 


