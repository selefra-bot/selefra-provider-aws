# Table: aws_ec2_instance_statuses

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| availability_zone | string | X | √ |  | 
| system_status | json | X | √ |  | 
| events | json | X | √ |  | 
| instance_id | string | X | √ |  | 
| instance_state | json | X | √ |  | 
| instance_status | json | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


