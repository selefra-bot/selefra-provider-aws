# Table: aws_accessanalyzer_analyzer_findings

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| is_public | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| analyzer_arn | string | X | √ |  | 
| analyzed_at | timestamp | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| error | string | X | √ |  | 
| arn | string | √ | √ |  | 
| condition | json | X | √ |  | 
| resource_type | string | X | √ |  | 
| action | string_array | X | √ |  | 
| resource | string | X | √ |  | 
| id | string | X | √ |  | 
| resource_owner_account | string | X | √ |  | 
| principal | json | X | √ |  | 
| sources | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_accessanalyzer_analyzers_selefra_id | string | X | X | fk to aws_accessanalyzer_analyzers.selefra_id | 


