# Table: aws_frauddetector_batch_predictions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| detector_version | string | X | √ |  | 
| detector_name | string | X | √ |  | 
| iam_role_arn | string | X | √ |  | 
| processed_records_count | big_int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| event_type_name | string | X | √ |  | 
| failure_reason | string | X | √ |  | 
| input_path | string | X | √ |  | 
| job_id | string | X | √ |  | 
| total_records_count | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| completion_time | string | X | √ |  | 
| last_heartbeat_time | string | X | √ |  | 
| output_path | string | X | √ |  | 
| start_time | string | X | √ |  | 
| status | string | X | √ |  | 


