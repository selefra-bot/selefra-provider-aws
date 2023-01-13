# Table: aws_glue_crawlers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| crawler_security_configuration | string | X | √ |  | 
| last_crawl | json | X | √ |  | 
| name | string | X | √ |  | 
| state | string | X | √ |  | 
| targets | json | X | √ |  | 
| crawl_elapsed_time | big_int | X | √ |  | 
| last_updated | timestamp | X | √ |  | 
| lineage_configuration | json | X | √ |  | 
| role | string | X | √ |  | 
| schedule | json | X | √ |  | 
| version | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| configuration | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| lake_formation_configuration | json | X | √ |  | 
| recrawl_policy | json | X | √ |  | 
| classifiers | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| database_name | string | X | √ |  | 
| description | string | X | √ |  | 
| schema_change_policy | json | X | √ |  | 
| table_prefix | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 


