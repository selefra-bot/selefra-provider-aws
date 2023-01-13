# Table: aws_ec2_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ena_support | bool | X | √ |  | 
| hypervisor | string | X | √ |  | 
| root_device_type | string | X | √ |  | 
| usage_operation_update_time | timestamp | X | √ |  | 
| image_id | string | X | √ |  | 
| platform_details | string | X | √ |  | 
| private_ip_address | string | X | √ |  | 
| sriov_net_support | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| enclave_options | json | X | √ |  | 
| launch_time | timestamp | X | √ |  | 
| tpm_support | string | X | √ |  | 
| capacity_reservation_specification | json | X | √ |  | 
| iam_instance_profile | json | X | √ |  | 
| kernel_id | string | X | √ |  | 
| state | json | X | √ |  | 
| region | string | X | √ |  | 
| instance_lifecycle | string | X | √ |  | 
| ramdisk_id | string | X | √ |  | 
| root_device_name | string | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| security_groups | json | X | √ |  | 
| source_dest_check | bool | X | √ |  | 
| usage_operation | string | X | √ |  | 
| ami_launch_index | big_int | X | √ |  | 
| elastic_gpu_associations | json | X | √ |  | 
| virtualization_type | string | X | √ |  | 
| elastic_inference_accelerator_associations | json | X | √ |  | 
| hibernation_options | json | X | √ |  | 
| state_reason | json | X | √ |  | 
| instance_type | string | X | √ |  | 
| ipv6_address | string | X | √ |  | 
| key_name | string | X | √ |  | 
| ebs_optimized | bool | X | √ |  | 
| instance_id | string | X | √ |  | 
| licenses | json | X | √ |  | 
| network_interfaces | json | X | √ |  | 
| private_dns_name_options | json | X | √ |  | 
| spot_instance_request_id | string | X | √ |  | 
| state_transition_reason | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| client_token | string | X | √ |  | 
| monitoring | json | X | √ |  | 
| platform | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| metadata_options | json | X | √ |  | 
| placement | json | X | √ |  | 
| public_dns_name | string | X | √ |  | 
| architecture | string | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| product_codes | json | X | √ |  | 
| public_ip_address | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| state_transition_reason_time | timestamp | X | √ |  | 
| capacity_reservation_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| block_device_mappings | json | X | √ |  | 
| boot_mode | string | X | √ |  | 
| cpu_options | json | X | √ |  | 
| maintenance_options | json | X | √ |  | 


