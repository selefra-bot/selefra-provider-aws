# Table: aws_qldb_ledgers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| state | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ | `The tags associated with the pipeline.` | 
| arn | string | X | √ |  | 
| creation_date_time | timestamp | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| encryption_description | json | X | √ |  | 
| permissions_mode | string | X | √ |  | 
| result_metadata | json | X | √ |  | 


