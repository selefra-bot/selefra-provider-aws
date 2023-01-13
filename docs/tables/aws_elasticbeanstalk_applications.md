# Table: aws_elasticbeanstalk_applications

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_lifecycle_config | json | X | √ |  | 
| versions | string_array | X | √ |  | 
| region | string | X | √ |  | 
| configuration_templates | string_array | X | √ |  | 
| date_updated | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| application_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| date_created | timestamp | X | √ |  | 
| application_arn | string | X | √ |  | 


