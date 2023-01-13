# Table: aws_ec2_instance_types

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| instance_type | string | X | √ |  | 
| placement_group_info | json | X | √ |  | 
| arn | string | √ | √ |  | 
| current_generation | bool | X | √ |  | 
| instance_storage_supported | bool | X | √ |  | 
| gpu_info | json | X | √ |  | 
| instance_storage_info | json | X | √ |  | 
| supported_boot_modes | string_array | X | √ |  | 
| v_cpu_info | json | X | √ |  | 
| burstable_performance_supported | bool | X | √ |  | 
| ebs_info | json | X | √ |  | 
| fpga_info | json | X | √ |  | 
| hypervisor | string | X | √ |  | 
| supported_virtualization_types | string_array | X | √ |  | 
| region | string | X | √ |  | 
| auto_recovery_supported | bool | X | √ |  | 
| hibernation_supported | bool | X | √ |  | 
| inference_accelerator_info | json | X | √ |  | 
| supported_root_device_types | string_array | X | √ |  | 
| supported_usage_classes | string_array | X | √ |  | 
| bare_metal | bool | X | √ |  | 
| dedicated_hosts_supported | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| free_tier_eligible | bool | X | √ |  | 
| network_info | json | X | √ |  | 
| memory_info | json | X | √ |  | 
| processor_info | json | X | √ |  | 


