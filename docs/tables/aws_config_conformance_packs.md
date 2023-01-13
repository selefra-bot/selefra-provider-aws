# Table: aws_config_conformance_packs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| conformance_pack_name | string | X | √ |  | 
| created_by | string | X | √ |  | 
| delivery_s3_bucket | string | X | √ |  | 
| template_ssm_document_details | json | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| conformance_pack_arn | string | X | √ |  | 
| conformance_pack_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| conformance_pack_input_parameters | json | X | √ |  | 
| delivery_s3_key_prefix | string | X | √ |  | 
| last_update_requested_time | timestamp | X | √ |  | 


