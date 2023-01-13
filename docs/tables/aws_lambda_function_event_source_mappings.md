# Table: aws_lambda_function_event_source_mappings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| maximum_record_age_in_seconds | big_int | X | √ |  | 
| maximum_retry_attempts | big_int | X | √ |  | 
| state_transition_reason | string | X | √ |  | 
| tumbling_window_in_seconds | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| bisect_batch_on_function_error | bool | X | √ |  | 
| maximum_batching_window_in_seconds | big_int | X | √ |  | 
| function_arn | string | X | √ |  | 
| function_response_types | string_array | X | √ |  | 
| last_processing_result | string | X | √ |  | 
| self_managed_kafka_event_source_config | json | X | √ |  | 
| amazon_managed_kafka_event_source_config | json | X | √ |  | 
| destination_config | json | X | √ |  | 
| starting_position_timestamp | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| uuid | string | X | √ |  | 
| filter_criteria | json | X | √ |  | 
| self_managed_event_source | json | X | √ |  | 
| source_access_configurations | json | X | √ |  | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| event_source_arn | string | X | √ |  | 
| parallelization_factor | big_int | X | √ |  | 
| queues | string_array | X | √ |  | 
| starting_position | string | X | √ |  | 
| topics | string_array | X | √ |  | 
| batch_size | big_int | X | √ |  | 
| last_modified | timestamp | X | √ |  | 


