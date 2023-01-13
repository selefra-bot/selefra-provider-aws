# Table: aws_elasticsearch_packages

## Primary Keys 

```
account_id, region, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| last_updated_at | timestamp | X | √ |  | 
| package_description | string | X | √ |  | 
| package_id | string | X | √ |  | 
| package_status | string | X | √ |  | 
| package_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| available_package_version | string | X | √ |  | 
| error_details | json | X | √ |  | 
| package_name | string | X | √ |  | 


