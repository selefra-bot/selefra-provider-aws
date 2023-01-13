# Table: aws_route53_domains

## Primary Keys 

```
account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| registrar_name | string | X | √ |  | 
| abuse_contact_phone | string | X | √ |  | 
| expiration_date | timestamp | X | √ |  | 
| who_is_server | string | X | √ |  | 
| account_id | string | √ | √ |  | 
| status_list | string_array | X | √ |  | 
| transfer_lock | bool | X | √ |  | 
| admin_privacy | bool | X | √ |  | 
| dns_sec | string | X | √ |  | 
| domain_name | string | X | √ |  | 
| tags | json | X | √ | `A list of tags` | 
| registrant_contact | json | X | √ |  | 
| registrar_url | string | X | √ |  | 
| abuse_contact_email | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| nameservers | json | X | √ |  | 
| reseller | string | X | √ |  | 
| tech_contact | json | X | √ |  | 
| dnssec_keys | json | X | √ |  | 
| registrant_privacy | bool | X | √ |  | 
| updated_date | timestamp | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| admin_contact | json | X | √ |  | 
| auto_renew | bool | X | √ |  | 
| registry_domain_id | string | X | √ |  | 
| tech_privacy | bool | X | √ |  | 


