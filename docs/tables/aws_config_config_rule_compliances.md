# Table: aws_config_config_rule_compliances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| compliance | json | X | √ |  | 
| config_rule_name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_config_config_rules_selefra_id | string | X | X | fk to aws_config_config_rules.selefra_id | 


