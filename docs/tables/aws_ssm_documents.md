# Table: aws_ssm_documents

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| platform_types | string_array | X | √ |  | 
| sha1 | string | X | √ |  | 
| requires | json | X | √ |  | 
| version_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| status_information | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| permissions | json | X | √ |  | 
| hash_type | string | X | √ |  | 
| latest_version | string | X | √ |  | 
| parameters | json | X | √ |  | 
| pending_review_version | string | X | √ |  | 
| attachments_information | json | X | √ |  | 
| review_status | string | X | √ |  | 
| status | string | X | √ |  | 
| schema_version | string | X | √ |  | 
| approved_version | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| default_version | string | X | √ |  | 
| description | string | X | √ |  | 
| display_name | string | X | √ |  | 
| document_version | string | X | √ |  | 
| name | string | X | √ |  | 
| region | string | X | √ |  | 
| author | string | X | √ |  | 
| category_enum | string_array | X | √ |  | 
| document_format | string | X | √ |  | 
| hash | string | X | √ |  | 
| category | string_array | X | √ |  | 
| document_type | string | X | √ |  | 
| owner | string | X | √ |  | 
| review_information | json | X | √ |  | 
| target_type | string | X | √ |  | 


