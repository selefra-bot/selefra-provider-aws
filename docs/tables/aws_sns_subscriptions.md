# Table: aws_sns_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| filter_policy | string | X | √ |  | 
| endpoint | string | X | √ |  | 
| owner | string | X | √ |  | 
| pending_confirmation | bool | X | √ |  | 
| unknown_fields | json | X | √ |  | 
| delivery_policy | string | X | √ |  | 
| effective_delivery_policy | string | X | √ |  | 
| protocol | string | X | √ |  | 
| subscription_arn | string | X | √ |  | 
| topic_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| redrive_policy | string | X | √ |  | 
| confirmation_was_authenticated | bool | X | √ |  | 
| arn | string | √ | √ |  | 
| raw_message_delivery | bool | X | √ |  | 
| subscription_role_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 


