# Table: aws_route53_traffic_policy_versions

## Primary Keys 

```
traffic_policy_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| traffic_policy_arn | string | √ | √ |  | 
| id | string | X | √ |  | 
| version | big_int | X | √ |  | 
| comment | string | X | √ |  | 
| document | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| aws_route53_traffic_policies_selefra_id | string | X | X | fk to aws_route53_traffic_policies.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


