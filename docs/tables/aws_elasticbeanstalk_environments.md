# Table: aws_elasticbeanstalk_environments

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| environment_name | string | X | √ |  | 
| tier | json | X | √ |  | 
| environment_id | string | X | √ |  | 
| solution_stack_name | string | X | √ |  | 
| template_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| abortable_operation_in_progress | bool | X | √ |  | 
| endpoint_url | string | X | √ |  | 
| listeners | json | X | √ |  | 
| cname | string | X | √ |  | 
| date_updated | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| version_label | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| health_status | string | X | √ |  | 
| platform_arn | string | X | √ |  | 
| resources | json | X | √ |  | 
| status | string | X | √ |  | 
| date_created | timestamp | X | √ |  | 
| environment_arn | string | X | √ |  | 
| environment_links | json | X | √ |  | 
| operations_role | string | X | √ |  | 
| tags | json | X | √ |  | 
| application_name | string | X | √ |  | 
| health | string | X | √ |  | 


