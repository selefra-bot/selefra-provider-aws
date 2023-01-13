# Table: aws_sagemaker_models

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ | `The tags associated with the model.` | 
| describe_model_output | json | X | √ |  | 
| model_arn | string | X | √ |  | 
| model_name | string | X | √ |  | 


