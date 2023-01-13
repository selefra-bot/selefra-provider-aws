# Table: aws_apigateway_rest_api_stages

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| canary_settings | json | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| last_updated_date | timestamp | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| access_log_settings | json | X | √ |  | 
| cache_cluster_status | string | X | √ |  | 
| variables | json | X | √ |  | 
| web_acl_arn | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| stage_name | string | X | √ |  | 
| tracing_enabled | bool | X | √ |  | 
| region | string | X | √ |  | 
| cache_cluster_enabled | bool | X | √ |  | 
| client_certificate_id | string | X | √ |  | 
| documentation_version | string | X | √ |  | 
| cache_cluster_size | string | X | √ |  | 
| method_settings | json | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| deployment_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


