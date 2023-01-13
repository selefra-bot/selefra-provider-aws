# Table: aws_lightsail_certificates

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subject_alternative_names | string_array | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| arn | string | X | √ |  | 
| in_use_resource_count | big_int | X | √ |  | 
| key_algorithm | string | X | √ |  | 
| name | string | X | √ |  | 
| revocation_reason | string | X | √ |  | 
| serial_number | string | X | √ |  | 
| tags | json | X | √ |  | 
| domain_name | string | X | √ |  | 
| issued_at | timestamp | X | √ |  | 
| issuer_ca | string | X | √ |  | 
| request_failure_reason | string | X | √ |  | 
| account_id | string | X | √ |  | 
| domain_validation_records | json | X | √ |  | 
| not_after | timestamp | X | √ |  | 
| renewal_summary | json | X | √ |  | 
| status | string | X | √ |  | 
| support_code | string | X | √ |  | 
| region | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| eligible_to_renew | string | X | √ |  | 
| not_before | timestamp | X | √ |  | 
| revoked_at | timestamp | X | √ |  | 


