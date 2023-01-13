# Table: aws_shield_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_arn | string | X | √ |  | 
| time_commitment_in_seconds | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| subscription_limits | json | X | √ |  | 
| auto_renew | string | X | √ |  | 
| end_time | timestamp | X | √ |  | 
| limits | json | X | √ |  | 
| arn | string | √ | √ |  | 
| proactive_engagement_status | string | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


