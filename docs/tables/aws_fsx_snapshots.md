# Table: aws_fsx_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| lifecycle | string | X | √ |  | 
| lifecycle_transition_reason | json | X | √ |  | 
| name | string | X | √ |  | 
| volume_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| resource_arn | string | X | √ |  | 
| snapshot_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| administrative_actions | json | X | √ |  | 


