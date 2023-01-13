# Table: aws_iam_group_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| group_id | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| aws_iam_groups_selefra_id | string | X | X | fk to aws_iam_groups.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| group_arn | string | X | √ |  | 
| policy_document | string | X | √ |  | 
| group_name | string | X | √ |  | 
| result_metadata | json | X | √ |  | 


