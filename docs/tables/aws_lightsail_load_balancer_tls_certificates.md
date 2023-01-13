# Table: aws_lightsail_load_balancer_tls_certificates

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| failure_reason | string | X | √ |  | 
| serial | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| load_balancer_name | string | X | √ |  | 
| location | json | X | √ |  | 
| resource_type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| tags | json | X | √ |  | 
| support_code | string | X | √ |  | 
| key_algorithm | string | X | √ |  | 
| is_attached | bool | X | √ |  | 
| not_before | timestamp | X | √ |  | 
| signature_algorithm | string | X | √ |  | 
| subject | string | X | √ |  | 
| domain_name | string | X | √ |  | 
| issued_at | timestamp | X | √ |  | 
| load_balancer_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| issuer | string | X | √ |  | 
| not_after | timestamp | X | √ |  | 
| renewal_summary | json | X | √ |  | 
| revocation_reason | string | X | √ |  | 
| subject_alternative_names | string_array | X | √ |  | 
| aws_lightsail_load_balancers_selefra_id | string | X | X | fk to aws_lightsail_load_balancers.selefra_id | 
| domain_validation_records | json | X | √ |  | 
| name | string | X | √ |  | 
| revoked_at | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 


