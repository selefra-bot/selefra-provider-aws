# Table: aws_ecr_repository_images

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| image_scan_findings_summary | json | X | √ |  | 
| image_scan_status | json | X | √ |  | 
| image_size_in_bytes | big_int | X | √ |  | 
| aws_ecr_repositories_selefra_id | string | X | X | fk to aws_ecr_repositories.selefra_id | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| registry_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| artifact_media_type | string | X | √ |  | 
| image_digest | string | X | √ |  | 
| image_pushed_at | timestamp | X | √ |  | 
| repository_name | string | X | √ |  | 
| image_manifest_media_type | string | X | √ |  | 
| image_tags | string_array | X | √ |  | 
| last_recorded_pull_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


