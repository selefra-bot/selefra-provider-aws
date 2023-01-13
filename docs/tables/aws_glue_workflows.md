# Table: aws_glue_workflows

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| default_run_properties | json | X | √ |  | 
| name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| blueprint_details | json | X | √ |  | 
| graph | json | X | √ |  | 
| last_run | json | X | √ |  | 
| arn | string | √ | √ |  | 
| last_modified_on | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| created_on | timestamp | X | √ |  | 
| max_concurrent_runs | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 


