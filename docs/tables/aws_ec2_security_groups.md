# Table: aws_ec2_security_groups

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
| description | string | X | √ |  | 
| group_name | string | X | √ |  | 
| ip_permissions | json | X | √ |  | 
| tags | json | X | √ |  | 
| group_id | string | X | √ |  | 
| ip_permissions_egress | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


