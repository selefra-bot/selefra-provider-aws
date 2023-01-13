# Table: aws_elbv2_listeners

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| alpn_policy | string_array | X | √ |  | 
| ssl_policy | string | X | √ |  | 
| region | string | X | √ |  | 
| default_actions | json | X | √ |  | 
| listener_arn | string | X | √ |  | 
| aws_elbv2_load_balancers_selefra_id | string | X | X | fk to aws_elbv2_load_balancers.selefra_id | 
| certificates | json | X | √ |  | 
| port | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| load_balancer_arn | string | X | √ |  | 
| protocol | string | X | √ |  | 


