# Table: aws_ssm_inventory_schemas

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| type_name | string | X | √ |  | 
| version | string | X | √ |  | 
| attributes | json | X | √ |  | 
| display_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


