# Table: aws_glue_ml_transforms

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| glue_version | string | X | √ |  | 
| number_of_workers | big_int | X | √ |  | 
| parameters | json | X | √ |  | 
| role | string | X | √ |  | 
| status | string | X | √ |  | 
| transform_id | string | X | √ |  | 
| worker_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| created_on | timestamp | X | √ |  | 
| evaluation_metrics | json | X | √ |  | 
| label_count | big_int | X | √ |  | 
| name | string | X | √ |  | 
| timeout | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| schema | json | X | √ |  | 
| last_modified_on | timestamp | X | √ |  | 
| max_capacity | float | X | √ |  | 
| description | string | X | √ |  | 
| input_record_tables | json | X | √ |  | 
| max_retries | big_int | X | √ |  | 
| transform_encryption | json | X | √ |  | 


