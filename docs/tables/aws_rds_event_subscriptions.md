# Table: aws_rds_event_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| source_ids_list | string_array | X | √ |  | 
| subscription_creation_time | string | X | √ |  | 
| tags | json | X | √ |  | 
| customer_aws_id | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| event_categories_list | string_array | X | √ |  | 
| sns_topic_arn | string | X | √ |  | 
| cust_subscription_id | string | X | √ |  | 
| event_subscription_arn | string | X | √ |  | 
| source_type | string | X | √ |  | 


