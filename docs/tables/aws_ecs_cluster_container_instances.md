# Table: aws_ecs_cluster_container_instances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| agent_connected | bool | X | √ |  | 
| attributes | json | X | √ |  | 
| ec2_instance_id | string | X | √ |  | 
| health_status | json | X | √ |  | 
| pending_tasks_count | int | X | √ |  | 
| registered_at | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| capacity_provider_name | string | X | √ |  | 
| registered_resources | json | X | √ |  | 
| remaining_resources | json | X | √ |  | 
| tags | json | X | √ |  | 
| version_info | json | X | √ |  | 
| container_instance_arn | string | X | √ |  | 
| status_reason | string | X | √ |  | 
| aws_ecs_clusters_selefra_id | string | X | X | fk to aws_ecs_clusters.selefra_id | 
| account_id | string | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| agent_update_status | string | X | √ |  | 
| attachments | json | X | √ |  | 
| running_tasks_count | int | X | √ |  | 
| version | int | X | √ |  | 


