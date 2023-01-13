# Table: aws_ssm_compliance_summary_items

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| compliance_type | string | X | √ |  | 
| compliant_summary | json | X | √ |  | 
| non_compliant_summary | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


