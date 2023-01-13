# Table: aws_ec2_transit_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| creation_time | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| options | json | X | √ |  | 
| state | string | X | √ |  | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| tags | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| transit_gateway_arn | string | X | √ |  | 
| transit_gateway_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


