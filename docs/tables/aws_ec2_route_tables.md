# Table: aws_ec2_route_tables

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| routes | json | X | √ |  | 
| vpc_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| associations | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| propagating_vgws | json | X | √ |  | 
| route_table_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


