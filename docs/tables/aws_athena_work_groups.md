# Table: aws_athena_work_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| configuration | json | X | √ |  | 
| tags | json | X | √ |  | 
| name | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 


