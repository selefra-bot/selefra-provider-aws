# Table: aws_dax_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| notification_configuration | json | X | √ |  | 
| parameter_group | json | X | √ |  | 
| status | string | X | √ |  | 
| subnet_group | string | X | √ |  | 
| total_nodes | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| node_ids_to_remove | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| cluster_discovery_endpoint | json | X | √ |  | 
| security_groups | json | X | √ |  | 
| description | string | X | √ |  | 
| iam_role_arn | string | X | √ |  | 
| node_type | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| cluster_endpoint_encryption_type | string | X | √ |  | 
| sse_description | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| active_nodes | big_int | X | √ |  | 
| cluster_name | string | X | √ |  | 
| nodes | json | X | √ |  | 


