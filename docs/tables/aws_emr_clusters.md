# Table: aws_emr_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| log_encryption_kms_key_id | string | X | √ |  | 
| step_concurrency_level | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ |  | 
| configurations | json | X | √ |  | 
| custom_ami_id | string | X | √ |  | 
| ec2_instance_attributes | json | X | √ |  | 
| applications | json | X | √ |  | 
| account_id | string | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| placement_groups | json | X | √ |  | 
| security_configuration | string | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| auto_terminate | bool | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| ebs_root_volume_size | big_int | X | √ |  | 
| os_release_label | string | X | √ |  | 
| release_label | string | X | √ |  | 
| running_ami_version | string | X | √ |  | 
| scale_down_behavior | string | X | √ |  | 
| auto_scaling_role | string | X | √ |  | 
| kerberos_attributes | json | X | √ |  | 
| service_role | string | X | √ |  | 
| status | json | X | √ |  | 
| log_uri | string | X | √ |  | 
| master_public_dns_name | string | X | √ |  | 
| requested_ami_version | string | X | √ |  | 
| termination_protected | bool | X | √ |  | 
| visible_to_all_users | bool | X | √ |  | 
| region | string | X | √ |  | 
| instance_collection_type | string | X | √ |  | 
| normalized_instance_hours | big_int | X | √ |  | 
| repo_upgrade_on_boot | string | X | √ |  | 


