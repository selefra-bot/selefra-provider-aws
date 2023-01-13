# Table: aws_ec2_vpn_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| amazon_side_asn | big_int | X | √ |  | 
| vpn_gateway_id | string | X | √ |  | 
| region | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| state | string | X | √ |  | 
| type | string | X | √ |  | 
| vpc_attachments | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


