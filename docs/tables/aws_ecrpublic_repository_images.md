# Table: aws_ecrpublic_repository_images

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| image_size_in_bytes | big_int | X | √ |  | 
| image_tags | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| image_manifest_media_type | string | X | √ |  | 
| image_pushed_at | timestamp | X | √ |  | 
| repository_name | string | X | √ |  | 
| aws_ecrpublic_repositories_selefra_id | string | X | X | fk to aws_ecrpublic_repositories.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| artifact_media_type | string | X | √ |  | 
| image_digest | string | X | √ |  | 
| registry_id | string | X | √ |  | 


