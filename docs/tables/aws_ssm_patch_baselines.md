# Table: aws_ssm_patch_baselines

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| baseline_description | string | X | √ |  | 
| baseline_name | string | X | √ |  | 
| default_baseline | bool | X | √ |  | 
| operating_system | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| baseline_id | string | X | √ |  | 


