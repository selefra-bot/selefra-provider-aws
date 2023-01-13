# Table: aws_s3_bucket_encryption_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| bucket_arn | string | X | √ |  | 
| apply_server_side_encryption_by_default | json | X | √ |  | 
| bucket_key_enabled | bool | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_s3_buckets_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| account_id | string | X | √ |  | 


