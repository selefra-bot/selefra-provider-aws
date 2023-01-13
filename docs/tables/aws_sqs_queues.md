# Table: aws_sqs_queues

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| url | string | X | √ |  | 
| policy | string | X | √ |  | 
| approximate_number_of_messages | big_int | X | √ |  | 
| sqs_managed_sse_enabled | bool | X | √ |  | 
| unknown_fields | json | X | √ |  | 
| created_timestamp | big_int | X | √ |  | 
| last_modified_timestamp | big_int | X | √ |  | 
| kms_data_key_reuse_period_seconds | big_int | X | √ |  | 
| approximate_number_of_messages_delayed | big_int | X | √ |  | 
| maximum_message_size | big_int | X | √ |  | 
| kms_master_key_id | string | X | √ |  | 
| content_based_deduplication | bool | X | √ |  | 
| deduplication_scope | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| redrive_allow_policy | string | X | √ |  | 
| tags | json | X | √ |  | 
| visibility_timeout | big_int | X | √ |  | 
| fifo_queue | bool | X | √ |  | 
| arn | string | X | √ |  | 
| redrive_policy | string | X | √ |  | 
| approximate_number_of_messages_not_visible | big_int | X | √ |  | 
| delay_seconds | big_int | X | √ |  | 
| fifo_throughput_limit | string | X | √ |  | 
| message_retention_period | big_int | X | √ |  | 
| receive_message_wait_time_seconds | big_int | X | √ |  | 


