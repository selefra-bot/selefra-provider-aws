# Table: aws_elasticbeanstalk_application_versions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| application_version_arn | string | X | √ |  | 
| date_created | timestamp | X | √ |  | 
| source_build_information | json | X | √ |  | 
| version_label | string | X | √ |  | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| build_arn | string | X | √ |  | 
| description | string | X | √ |  | 
| source_bundle | json | X | √ |  | 
| region | string | X | √ |  | 
| application_name | string | X | √ |  | 
| date_updated | timestamp | X | √ |  | 


