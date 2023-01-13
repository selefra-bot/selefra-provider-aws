# Table: aws_docdb_cluster_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| attributes | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| percent_progress | big_int | X | √ |  | 
| aws_docdb_clusters_selefra_id | string | X | X | fk to aws_docdb_clusters.selefra_id | 
| arn | string | √ | √ |  | 
| db_cluster_snapshot_identifier | string | X | √ |  | 
| db_cluster_snapshot_arn | string | X | √ |  | 
| engine | string | X | √ |  | 
| port | big_int | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| engine_version | string | X | √ |  | 
| master_username | string | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| source_db_cluster_snapshot_arn | string | X | √ |  | 
| vpc_id | string | X | √ |  | 


