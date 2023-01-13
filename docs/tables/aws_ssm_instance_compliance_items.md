# Table: aws_ssm_instance_compliance_items

## Primary Keys 

```
instance_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_type | string | X | √ |  | 
| aws_ssm_instances_selefra_id | string | X | X | fk to aws_ssm_instances.selefra_id | 
| account_id | string | X | √ |  | 
| instance_arn | string | √ | √ |  | 
| compliance_type | string | X | √ |  | 
| execution_summary | json | X | √ |  | 
| title | string | X | √ |  | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| details | json | X | √ |  | 
| resource_id | string | X | √ |  | 
| severity | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


