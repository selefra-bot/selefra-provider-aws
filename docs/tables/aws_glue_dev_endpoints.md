# Table: aws_glue_dev_endpoints

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| security_group_ids | string_array | X | √ |  | 
| zeppelin_remote_spark_interpreter_port | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| public_key | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| public_keys | string_array | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arguments | json | X | √ |  | 
| failure_reason | string | X | √ |  | 
| last_modified_timestamp | timestamp | X | √ |  | 
| private_address | string | X | √ |  | 
| last_update_status | string | X | √ |  | 
| worker_type | string | X | √ |  | 
| yarn_endpoint_address | string | X | √ |  | 
| region | string | X | √ |  | 
| created_timestamp | timestamp | X | √ |  | 
| number_of_workers | big_int | X | √ |  | 
| role_arn | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| number_of_nodes | big_int | X | √ |  | 
| public_address | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| security_configuration | string | X | √ |  | 
| arn | string | √ | √ |  | 
| endpoint_name | string | X | √ |  | 
| extra_jars_s3_path | string | X | √ |  | 
| extra_python_libs_s3_path | string | X | √ |  | 
| glue_version | string | X | √ |  | 


