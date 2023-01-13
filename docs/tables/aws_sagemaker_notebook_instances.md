# Table: aws_sagemaker_notebook_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ | `The tags associated with the notebook instance.` | 
| describe_notebook_instance_output | json | X | √ |  | 
| notebook_instance_arn | string | X | √ |  | 
| notebook_instance_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


