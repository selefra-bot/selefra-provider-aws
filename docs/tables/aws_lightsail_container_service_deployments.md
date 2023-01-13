# Table: aws_lightsail_container_service_deployments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_lightsail_container_services_selefra_id | string | X | X | fk to aws_lightsail_container_services.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| container_service_arn | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| containers | json | X | √ |  | 
| public_endpoint | json | X | √ |  | 
| state | string | X | √ |  | 
| version | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


