# Table: aws_ram_resource_share_invitations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_share_arn | string | X | √ |  | 
| resource_share_associations | json | X | √ |  | 
| resource_share_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| invitation_timestamp | timestamp | X | √ |  | 
| resource_share_invitation_arn | string | X | √ |  | 
| sender_account_id | string | X | √ |  | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| receiver_account_id | string | X | √ |  | 
| receiver_arn | string | X | √ |  | 


