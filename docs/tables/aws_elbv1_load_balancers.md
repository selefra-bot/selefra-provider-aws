# Table: aws_elbv1_load_balancers

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
| load_balancer_description | json | X | √ |  | 
| tags | json | X | √ |  | 
| attributes | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


