# Table: aws_directconnect_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| amazon_side_asn | big_int | X | √ |  | 
| direct_connect_gateway_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| id | string | X | √ |  | 
| direct_connect_gateway_name | string | X | √ |  | 
| direct_connect_gateway_state | string | X | √ |  | 
| owner_account | string | X | √ |  | 
| state_change_error | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


