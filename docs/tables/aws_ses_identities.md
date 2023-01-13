# Table: aws_ses_identities

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| identity_name | string | X | √ |  | 
| sending_enabled | bool | X | √ |  | 
| get_email_identity_output | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


