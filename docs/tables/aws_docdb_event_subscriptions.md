# Table: aws_docdb_event_subscriptions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| customer_aws_id | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| event_subscription_arn | string | X | √ |  | 
| source_type | string | X | √ |  | 
| subscription_creation_time | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| cust_subscription_id | string | X | √ |  | 
| event_categories_list | string_array | X | √ |  | 
| sns_topic_arn | string | X | √ |  | 
| source_ids_list | string_array | X | √ |  | 
| status | string | X | √ |  | 


