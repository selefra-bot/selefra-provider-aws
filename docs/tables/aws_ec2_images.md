# Table: aws_ec2_images

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| kernel_id | string | X | √ |  | 
| platform_details | string | X | √ |  | 
| root_device_name | string | X | √ |  | 
| description | string | X | √ |  | 
| virtualization_type | string | X | √ |  | 
| image_location | string | X | √ |  | 
| root_device_type | string | X | √ |  | 
| sriov_net_support | string | X | √ |  | 
| state_reason | json | X | √ |  | 
| tags | json | X | √ |  | 
| creation_date | string | X | √ |  | 
| ena_support | bool | X | √ |  | 
| product_codes | json | X | √ |  | 
| architecture | string | X | √ |  | 
| image_owner_alias | string | X | √ |  | 
| image_type | string | X | √ |  | 
| imds_support | string | X | √ |  | 
| state | string | X | √ |  | 
| account_id | string | X | √ |  | 
| block_device_mappings | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| deprecation_time | string | X | √ |  | 
| hypervisor | string | X | √ |  | 
| image_id | string | X | √ |  | 
| public | bool | X | √ |  | 
| tpm_support | string | X | √ |  | 
| boot_mode | string | X | √ |  | 
| name | string | X | √ |  | 
| platform | string | X | √ |  | 
| ramdisk_id | string | X | √ |  | 
| usage_operation | string | X | √ |  | 


