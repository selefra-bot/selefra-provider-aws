# Table: aws_sagemaker_endpoint_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| endpoint_config_arn | string | X | √ |  | 
| endpoint_config_name | string | X | √ |  | 
| async_inference_config | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| data_capture_config | json | X | √ |  | 
| explainer_config | json | X | √ |  | 
| shadow_production_variants | json | X | √ |  | 
| account_id | string | X | √ |  | 
| production_variants | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ | `The tags associated with the model.` | 
| creation_time | timestamp | X | √ |  | 


