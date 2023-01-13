# Table: aws_ec2_nat_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| state | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| delete_time | timestamp | X | √ |  | 
| provisioned_bandwidth | json | X | √ |  | 
| failure_message | string | X | √ |  | 
| nat_gateway_addresses | json | X | √ |  | 
| account_id | string | X | √ |  | 
| failure_code | string | X | √ |  | 
| connectivity_type | string | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| nat_gateway_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 


