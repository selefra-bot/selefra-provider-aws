# Table: aws_ses_configuration_set_event_destinations

## Primary Keys 

```
account_id, region, configuration_set_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| matching_event_types | string_array | X | √ |  | 
| kinesis_firehose_destination | json | X | √ |  | 
| pinpoint_destination | json | X | √ |  | 
| aws_ses_configuration_sets_selefra_id | string | X | X | fk to aws_ses_configuration_sets.selefra_id | 
| name | string | X | √ |  | 
| region | string | X | √ |  | 
| configuration_set_name | string | X | √ |  | 
| cloud_watch_destination | json | X | √ |  | 
| enabled | bool | X | √ |  | 
| sns_destination | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


