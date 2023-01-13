# Table: aws_ses_contact_lists

## Primary Keys 

```
account_id, region, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| contact_list_name | string | X | √ |  | 
| created_timestamp | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| last_updated_timestamp | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| topics | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


