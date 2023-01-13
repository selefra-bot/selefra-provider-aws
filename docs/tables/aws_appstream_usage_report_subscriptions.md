# Table: aws_appstream_usage_report_subscriptions

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| s3_bucket_name | string | X | √ |  | 
| last_generated_report_date | timestamp | X | √ |  | 
| schedule | string | X | √ |  | 
| subscription_errors | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


