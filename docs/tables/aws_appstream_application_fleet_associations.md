# Table: aws_appstream_application_fleet_associations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_appstream_applications_selefra_id | string | X | X | fk to aws_appstream_applications.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| application_arn | string | X | √ |  | 
| fleet_name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


