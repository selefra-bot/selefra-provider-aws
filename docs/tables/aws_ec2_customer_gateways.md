# Table: aws_ec2_customer_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| customer_gateway_id | string | X | √ |  | 
| device_name | string | X | √ |  | 
| ip_address | string | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| bgp_asn | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| certificate_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| type | string | X | √ |  | 
| account_id | string | X | √ |  | 


