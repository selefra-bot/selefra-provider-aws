# Table: aws_ec2_key_pairs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| create_time | timestamp | X | √ |  | 
| key_fingerprint | string | X | √ |  | 
| key_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| key_pair_id | string | X | √ |  | 
| key_type | string | X | √ |  | 
| public_key | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 


