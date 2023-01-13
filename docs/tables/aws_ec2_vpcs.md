# Table: aws_ec2_vpcs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| cidr_block_association_set | json | X | √ |  | 
| dhcp_options_id | string | X | √ |  | 
| state | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| instance_tenancy | string | X | √ |  | 
| arn | string | √ | √ |  | 
| is_default | bool | X | √ |  | 
| owner_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| cidr_block | string | X | √ |  | 
| ipv6_cidr_block_association_set | json | X | √ |  | 


