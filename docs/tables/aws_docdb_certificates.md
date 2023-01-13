# Table: aws_docdb_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| valid_from | timestamp | X | √ |  | 
| valid_till | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| certificate_type | string | X | √ |  | 
| thumbprint | string | X | √ |  | 
| certificate_identifier | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| certificate_arn | string | X | √ |  | 


