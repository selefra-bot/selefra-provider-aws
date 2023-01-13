# Table: aws_inspector2_findings

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_account_id | string | X | √ |  | 
| inspector_score | float | X | √ |  | 
| account_id | string | X | √ |  | 
| finding_arn | string | X | √ |  | 
| remediation | json | X | √ |  | 
| resources | json | X | √ |  | 
| type | string | X | √ |  | 
| inspector_score_details | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| description | string | X | √ |  | 
| status | string | X | √ |  | 
| fix_available | string | X | √ |  | 
| network_reachability_details | json | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| first_observed_at | timestamp | X | √ |  | 
| last_observed_at | timestamp | X | √ |  | 
| severity | string | X | √ |  | 
| exploit_available | string | X | √ |  | 
| exploitability_details | json | X | √ |  | 
| package_vulnerability_details | json | X | √ |  | 
| title | string | X | √ |  | 


