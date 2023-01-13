# Table: aws_elasticbeanstalk_configuration_options

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| environment_id | string | X | √ |  | 
| configuration_option_description | json | X | √ |  | 
| application_arn | string | X | √ |  | 
| aws_elasticbeanstalk_environments_selefra_id | string | X | X | fk to aws_elasticbeanstalk_environments.selefra_id | 
| selefra_id | string | √ | √ | random id | 


