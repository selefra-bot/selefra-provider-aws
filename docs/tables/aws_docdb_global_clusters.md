# Table: aws_docdb_global_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| global_cluster_identifier | string | X | √ |  | 
| global_cluster_members | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| deletion_protection | bool | X | √ |  | 
| engine | string | X | √ |  | 
| global_cluster_resource_id | string | X | √ |  | 
| status | string | X | √ |  | 
| database_name | string | X | √ |  | 
| global_cluster_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 


