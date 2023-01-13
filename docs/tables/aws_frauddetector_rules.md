# Table: aws_frauddetector_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| detector_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| expression | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| arn | string | X | √ |  | 
| description | string | X | √ |  | 
| last_updated_time | string | X | √ |  | 
| rule_version | string | X | √ |  | 
| aws_frauddetector_detectors_selefra_id | string | X | X | fk to aws_frauddetector_detectors.selefra_id | 
| region | string | X | √ |  | 
| created_time | string | X | √ |  | 
| language | string | X | √ |  | 
| outcomes | string_array | X | √ |  | 
| rule_id | string | X | √ |  | 


