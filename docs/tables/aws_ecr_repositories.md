# Table: aws_ecr_repositories

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| policy_text | json | X | √ |  | 
| image_tag_mutability | string | X | √ |  | 
| repository_name | string | X | √ |  | 
| repository_uri | string | X | √ |  | 
| repository_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| registry_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| encryption_configuration | json | X | √ |  | 
| image_scanning_configuration | json | X | √ |  | 


