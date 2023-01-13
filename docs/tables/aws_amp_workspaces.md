# Table: aws_amp_workspaces

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| logging_configuration | json | X | √ |  | 
| arn | string | X | √ |  | 
| status | json | X | √ |  | 
| workspace_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| alert_manager_definition | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| alias | string | X | √ |  | 
| prometheus_endpoint | string | X | √ |  | 
| tags | json | X | √ |  | 


