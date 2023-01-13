# Table: aws_qldb_ledger_journal_kinesis_streams

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ledger_name | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| stream_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_qldb_ledgers_selefra_id | string | X | X | fk to aws_qldb_ledgers.selefra_id | 
| ledger_arn | string | X | √ |  | 
| kinesis_configuration | json | X | √ |  | 
| inclusive_start_time | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| exclusive_end_time | timestamp | X | √ |  | 
| error_cause | string | X | √ |  | 
| arn | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| stream_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| status | string | X | √ |  | 


