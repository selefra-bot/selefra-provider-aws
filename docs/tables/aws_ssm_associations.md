# Table: aws_ssm_associations

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_execution_date | timestamp | X | √ |  | 
| overview | json | X | √ |  | 
| targets | json | X | √ |  | 
| account_id | string | X | √ |  | 
| association_id | string | X | √ |  | 
| instance_id | string | X | √ |  | 
| schedule_expression | string | X | √ |  | 
| schedule_offset | big_int | X | √ |  | 
| target_maps | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| association_name | string | X | √ |  | 
| name | string | X | √ |  | 
| association_version | string | X | √ |  | 
| document_version | string | X | √ |  | 


