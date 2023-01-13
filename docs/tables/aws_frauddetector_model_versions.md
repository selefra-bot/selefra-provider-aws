# Table: aws_frauddetector_model_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| external_events_detail | json | X | √ |  | 
| ingested_events_detail | json | X | √ |  | 
| status | string | X | √ |  | 
| created_time | string | X | √ |  | 
| model_type | string | X | √ |  | 
| model_version_number | string | X | √ |  | 
| training_data_schema | json | X | √ |  | 
| training_result | json | X | √ |  | 
| training_result_v2 | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| model_id | string | X | √ |  | 
| training_data_source | string | X | √ |  | 
| aws_frauddetector_models_selefra_id | string | X | X | fk to aws_frauddetector_selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| last_updated_time | string | X | √ |  | 


