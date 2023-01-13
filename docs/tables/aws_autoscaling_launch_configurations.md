# Table: aws_autoscaling_launch_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| image_id | string | X | √ |  | 
| block_device_mappings | json | X | √ |  | 
| iam_instance_profile | string | X | √ |  | 
| metadata_options | json | X | √ |  | 
| user_data | string | X | √ |  | 
| ebs_optimized | bool | X | √ |  | 
| instance_monitoring | json | X | √ |  | 
| kernel_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| launch_configuration_name | string | X | √ |  | 
| associate_public_ip_address | bool | X | √ |  | 
| classic_link_vpc_id | string | X | √ |  | 
| key_name | string | X | √ |  | 
| placement_tenancy | string | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| spot_price | string | X | √ |  | 
| instance_type | string | X | √ |  | 
| classic_link_vpc_security_groups | string_array | X | √ |  | 
| launch_configuration_arn | string | X | √ |  | 
| ramdisk_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


