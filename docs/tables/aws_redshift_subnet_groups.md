# Table: aws_redshift_subnet_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ | `The Amazon Resource Name (ARN) for the resource.` | 
| tags | json | X | √ |  | 
| cluster_subnet_group_name | string | X | √ |  | 
| subnets | json | X | √ |  | 
| vpc_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| subnet_group_status | string | X | √ |  | 


