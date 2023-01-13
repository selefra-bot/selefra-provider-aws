# Table: aws_ssm_instance_patches

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| classification | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_ssm_instances_selefra_id | string | X | X | fk to aws_ssm_instances.selefra_id | 
| account_id | string | X | √ |  | 
| kb_id | string | X | √ |  | 
| severity | string | X | √ |  | 
| state | string | X | √ |  | 
| title | string | X | √ |  | 
| cve_ids | string | X | √ |  | 
| region | string | X | √ |  | 
| installed_time | timestamp | X | √ |  | 


