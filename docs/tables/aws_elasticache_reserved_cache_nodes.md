# Table: aws_elasticache_reserved_cache_nodes

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cache_node_count | big_int | X | √ |  | 
| product_description | string | X | √ |  | 
| usage_price | float | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| reserved_cache_nodes_offering_id | string | X | √ |  | 
| fixed_price | float | X | √ |  | 
| offering_type | string | X | √ |  | 
| recurring_charges | json | X | √ |  | 
| reservation_arn | string | X | √ |  | 
| reserved_cache_node_id | string | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| duration | big_int | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


