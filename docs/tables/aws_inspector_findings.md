# Table: aws_inspector_findings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| confidence | big_int | X | √ |  | 
| schema_version | big_int | X | √ |  | 
| title | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| severity | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| id | string | X | √ |  | 
| indicator_of_compromise | bool | X | √ |  | 
| recommendation | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| attributes | json | X | √ |  | 
| user_attributes | json | X | √ |  | 
| asset_type | string | X | √ |  | 
| service_attributes | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| asset_attributes | json | X | √ |  | 
| description | string | X | √ |  | 
| numeric_severity | float | X | √ |  | 
| service | string | X | √ |  | 


