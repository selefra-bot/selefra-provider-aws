# Table: aws_neptune_global_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| engine | string | X | √ |  | 
| global_cluster_identifier | string | X | √ |  | 
| account_id | string | X | √ |  | 
| global_cluster_resource_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| engine_version | string | X | √ |  | 
| global_cluster_members | json | X | √ |  | 
| status | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| region | string | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| global_cluster_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


