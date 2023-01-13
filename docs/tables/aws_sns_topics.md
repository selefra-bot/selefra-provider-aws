# Table: aws_sns_topics

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| delivery_policy | string | X | √ |  | 
| policy | string | X | √ |  | 
| owner | string | X | √ |  | 
| effective_delivery_policy | string | X | √ |  | 
| display_name | string | X | √ |  | 
| subscriptions_pending | big_int | X | √ |  | 
| kms_master_key_id | string | X | √ |  | 
| content_based_deduplication | bool | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| subscriptions_deleted | big_int | X | √ |  | 
| unknown_fields | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| subscriptions_confirmed | big_int | X | √ |  | 
| fifo_topic | bool | X | √ |  | 


