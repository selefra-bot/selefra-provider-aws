# Table: aws_rds_engine_versions

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| engine | string | X | √ |  | 
| db_engine_version_arn | string | X | √ |  | 
| image | json | X | √ |  | 
| supports_log_exports_to_cloudwatch_logs | bool | X | √ |  | 
| region | string | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| custom_db_engine_version_manifest | string | X | √ |  | 
| db_parameter_group_family | string | X | √ |  | 
| default_character_set | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| supports_certificate_rotation_without_restart | bool | X | √ |  | 
| db_engine_version_description | string | X | √ |  | 
| database_installation_files_s3_prefix | string | X | √ |  | 
| supported_character_sets | json | X | √ |  | 
| valid_upgrade_target | json | X | √ |  | 
| account_id | string | X | √ |  | 
| tag_list | json | X | √ |  | 
| supports_read_replica | bool | X | √ |  | 
| supported_engine_modes | string_array | X | √ |  | 
| supported_feature_names | string_array | X | √ |  | 
| supported_nchar_character_sets | json | X | √ |  | 
| supports_babelfish | bool | X | √ |  | 
| engine_version | string | X | √ |  | 
| exportable_log_types | string_array | X | √ |  | 
| status | string | X | √ |  | 
| supported_ca_certificate_identifiers | string_array | X | √ |  | 
| supported_timezones | json | X | √ |  | 
| supports_global_databases | bool | X | √ |  | 
| supports_parallel_query | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| db_engine_description | string | X | √ |  | 
| db_engine_media_type | string | X | √ |  | 
| database_installation_files_s3_bucket_name | string | X | √ |  | 
| major_engine_version | string | X | √ |  | 


