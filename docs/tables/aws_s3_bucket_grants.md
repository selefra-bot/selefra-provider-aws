# Table: aws_s3_bucket_grants

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| bucket_arn | string | X | √ |  | 
| grantee | json | X | √ |  | 
| permission | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_s3_buckets_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 


