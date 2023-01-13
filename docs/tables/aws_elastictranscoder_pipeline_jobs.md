# Table: aws_elastictranscoder_pipeline_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| input | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| output_key_prefix | string | X | √ |  | 
| outputs | json | X | √ |  | 
| playlists | json | X | √ |  | 
| aws_elastictranscoder_pipelines_selefra_id | string | X | X | fk to aws_elastictranscoder_pipelines.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| pipeline_id | string | X | √ |  | 
| user_metadata | json | X | √ |  | 
| arn | string | X | √ |  | 
| inputs | json | X | √ |  | 
| output | json | X | √ |  | 
| status | string | X | √ |  | 
| timing | json | X | √ |  | 


