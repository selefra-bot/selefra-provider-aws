# Table: aws_guardduty_detector_members

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| email | string | X | √ |  | 
| relationship_status | string | X | √ |  | 
| updated_at | string | X | √ |  | 
| administrator_id | string | X | √ |  | 
| detector_id | string | X | √ |  | 
| invited_at | string | X | √ |  | 
| region | string | X | √ |  | 
| detector_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_guardduty_detectors_selefra_id | string | X | X | fk to aws_guardduty_detectors.selefra_id | 
| account_id | string | X | √ |  | 
| master_id | string | X | √ |  | 


