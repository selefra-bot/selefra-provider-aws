# Table: aws_ec2_reserved_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| duration | big_int | X | √ |  | 
| end | timestamp | X | √ |  | 
| fixed_price | float | X | √ |  | 
| offering_class | string | X | √ |  | 
| product_description | string | X | √ |  | 
| usage_price | float | X | √ |  | 
| reserved_instances_id | string | X | √ |  | 
| start | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| availability_zone | string | X | √ |  | 
| currency_code | string | X | √ |  | 
| instance_count | big_int | X | √ |  | 
| instance_tenancy | string | X | √ |  | 
| instance_type | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| offering_type | string | X | √ |  | 
| recurring_charges | json | X | √ |  | 
| scope | string | X | √ |  | 


