# Table: aws_iam_virtual_mfa_devices

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| serial_number | string | X | √ |  | 
| tags | json | X | √ |  | 
| base32_string_seed | int_array | X | √ |  | 
| enable_date | timestamp | X | √ |  | 
| qr_code_png | int_array | X | √ |  | 
| user | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 


