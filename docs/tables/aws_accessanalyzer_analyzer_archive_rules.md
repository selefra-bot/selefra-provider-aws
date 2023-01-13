# Table: aws_accessanalyzer_analyzer_archive_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| analyzer_arn | string | X | √ |  | 
| filter | json | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| aws_accessanalyzer_analyzers_selefra_id | string | X | X | fk to aws_accessanalyzer_analyzers.selefra_id | 
| account_id | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| rule_name | string | X | √ |  | 


