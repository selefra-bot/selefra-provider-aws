# Table: aws_redshift_event_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ | `ARN of the event subscription.` | 
| tags | json | X | √ |  | 
| cust_subscription_id | string | X | √ |  | 
| customer_aws_id | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| region | string | X | √ |  | 
| source_ids_list | string_array | X | √ |  | 
| source_type | string | X | √ |  | 
| status | string | X | √ |  | 
| severity | string | X | √ |  | 
| sns_topic_arn | string | X | √ |  | 
| subscription_creation_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| event_categories_list | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


