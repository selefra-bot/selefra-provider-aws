# Table: aws_config_conformance_pack_rule_compliances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_config_conformance_packs_selefra_id | string | X | X | fk to aws_config_conformance_packs.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| conformance_pack_arn | string | X | √ |  | 
| compliance_type | string | X | √ |  | 
| evaluation_result_identifier | json | X | √ |  | 
| config_rule_name | string | X | √ |  | 
| controls | string_array | X | √ |  | 
| config_rule_invoked_time | timestamp | X | √ |  | 
| result_recorded_time | timestamp | X | √ |  | 
| annotation | string | X | √ |  | 


