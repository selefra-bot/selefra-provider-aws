# Table: aws_cloudwatchlogs_resource_policies

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| policy_document | string | X | √ |  | 
| last_updated_time | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


