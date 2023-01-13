# Table: aws_ses_custom_verification_email_templates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| template_content | string | X | √ |  | 
| template_subject | string | X | √ |  | 
| failure_redirection_url | string | X | √ |  | 
| success_redirection_url | string | X | √ |  | 
| arn | string | √ | √ |  | 
| from_email_address | string | X | √ |  | 
| template_name | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


