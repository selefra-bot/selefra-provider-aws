# Table: aws_ec2_vpc_peering_connections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| accepter_vpc_info | json | X | √ |  | 
| expiration_time | timestamp | X | √ |  | 
| requester_vpc_info | json | X | √ |  | 
| status | json | X | √ |  | 
| vpc_peering_connection_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 


