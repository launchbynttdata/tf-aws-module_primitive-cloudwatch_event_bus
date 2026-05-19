# Complete Example

This example creates an EventBridge event bus with customer-managed KMS encryption, using the resource naming module for consistent naming.

## Usage

```hcl
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

module "resource_names" {
  source  = "terraform.registry.launch.nttdata.com/module_library/resource_name/launch"
  version = "~> 2.0"

  for_each = var.resource_names_map

  logical_product_family  = var.logical_product_family
  logical_product_service = var.logical_product_service
  class_env               = var.class_env
  instance_env            = var.instance_env
  instance_resource       = var.instance_resource
  cloud_resource_type     = each.value.name
  maximum_length          = each.value.max_length
  region                  = join("", split("-", data.aws_region.current.name))
}

resource "aws_kms_key" "event_bus" {
  description             = "KMS key for EventBridge event bus ${module.resource_names["eventbus"].standard}"
  deletion_window_in_days  = 7
  enable_key_rotation     = true

  tags = var.tags
}

resource "aws_kms_key_policy" "event_bus" {
  key_id = aws_kms_key.event_bus.id
  policy = jsonencode({
    Version = "2012-10-17"
    Id      = "eventbridge-key-policy"
    Statement = [
      {
        Sid    = "Enable IAM User Permissions"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action   = "kms:*"
        Resource = "*"
      },
      {
        Sid    = "Allow EventBridge to use the key"
        Effect = "Allow"
        Principal = {
          Service = "events.amazonaws.com"
        }
        Action = [
          "kms:GenerateDataKey",
          "kms:Decrypt",
          "kms:DescribeKey"
        ]
        Resource = "*"
        Condition = {
          StringEquals = {
            "aws:SourceAccount" = data.aws_caller_identity.current.account_id
          }
        }
      }
    ]
  })
}

resource "aws_kms_alias" "event_bus" {
  name          = "alias/event-bus-${module.resource_names["eventbus"].standard}"
  target_key_id = aws_kms_key.event_bus.key_id
}

module "event_bus" {
  source = "../.."

  name               = module.resource_names["eventbus"].standard
  description        = var.description
  event_source_name  = var.event_source_name
  kms_key_identifier = aws_kms_key.event_bus.arn
  dead_letter_config = var.dead_letter_config
  tags               = var.tags

  depends_on = [aws_kms_key_policy.event_bus]
}
```

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_logical_product_family"></a> [logical_product_family](#input_logical_product_family) | Logical product family for resource naming. | `string` | n/a | yes |
| <a name="input_logical_product_service"></a> [logical_product_service](#input_logical_product_service) | Logical product service for resource naming. | `string` | n/a | yes |
| <a name="input_class_env"></a> [class_env](#input_class_env) | Class environment for resource naming (e.g., dev, prod). | `string` | n/a | yes |
| <a name="input_instance_env"></a> [instance_env](#input_instance_env) | Instance environment for resource naming. | `string` | n/a | yes |
| <a name="input_instance_resource"></a> [instance_resource](#input_instance_resource) | Instance resource for resource naming (numeric identifier). | `number` | n/a | yes |
| <a name="input_resource_names_map"></a> [resource_names_map](#input_resource_names_map) | Map of resource types to naming configuration (name must be letters and numbers only; max_length). | `map(object({ name = string, max_length = number }))` | n/a | yes |
| <a name="input_description"></a> [description](#input_description) | Event bus description. | `string` | `null` | no |
| <a name="input_event_source_name"></a> [event_source_name](#input_event_source_name) | Partner event source that the new event bus will be matched with. Must match name for partner event buses. | `string` | `null` | no |
| <a name="input_dead_letter_config"></a> [dead_letter_config](#input_dead_letter_config) | Configuration details of the Amazon SQS queue for EventBridge to use as a dead-letter queue (DLQ). | `object({ arn = optional(string) })` | `null` | no |
| <a name="input_tags"></a> [tags](#input_tags) | Map of tags to assign to the event bus. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output_id) | The ID of the event bus (same as the name). |
| <a name="output_arn"></a> [arn](#output_arn) | The ARN of the event bus. |
| <a name="output_name"></a> [name](#output_name) | The name of the event bus. |
| <a name="output_kms_key_arn"></a> [kms_key_arn](#output_kms_key_arn) | The ARN of the KMS key used for event bus encryption. |

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.14 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_event_bus"></a> [event\_bus](#module\_event\_bus) | ../.. | n/a |
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | terraform.registry.launch.nttdata.com/module_library/resource_name/launch | ~> 2.0 |

## Resources

| Name | Type |
|------|------|
| [aws_kms_alias.event_bus](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_alias) | resource |
| [aws_kms_key.event_bus](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key) | resource |
| [aws_kms_key_policy.event_bus](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key_policy) | resource |
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_class_env"></a> [class\_env](#input\_class\_env) | Class environment for resource naming (e.g., dev, prod). | `string` | n/a | yes |
| <a name="input_dead_letter_config"></a> [dead\_letter\_config](#input\_dead\_letter\_config) | Configuration details of the Amazon SQS queue for EventBridge to use as a dead-letter queue (DLQ). | <pre>object({<br/>    arn = optional(string)<br/>  })</pre> | `null` | no |
| <a name="input_description"></a> [description](#input\_description) | Event bus description. | `string` | `null` | no |
| <a name="input_event_source_name"></a> [event\_source\_name](#input\_event\_source\_name) | Partner event source that the new event bus will be matched with. Must match name for partner event buses. | `string` | `null` | no |
| <a name="input_instance_env"></a> [instance\_env](#input\_instance\_env) | Instance environment for resource naming. | `string` | n/a | yes |
| <a name="input_instance_resource"></a> [instance\_resource](#input\_instance\_resource) | Instance resource for resource naming (numeric identifier). | `number` | n/a | yes |
| <a name="input_logical_product_family"></a> [logical\_product\_family](#input\_logical\_product\_family) | Logical product family for resource naming. | `string` | n/a | yes |
| <a name="input_logical_product_service"></a> [logical\_product\_service](#input\_logical\_product\_service) | Logical product service for resource naming. | `string` | n/a | yes |
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | Map of resource types to naming configuration (name, max\_length). | <pre>map(object({<br/>    name       = string<br/>    max_length = number<br/>  }))</pre> | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | Map of tags to assign to the event bus. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the event bus. |
| <a name="output_id"></a> [id](#output\_id) | The ID of the event bus (same as the name). |
| <a name="output_kms_key_arn"></a> [kms\_key\_arn](#output\_kms\_key\_arn) | The ARN of the KMS key used for event bus encryption. |
| <a name="output_name"></a> [name](#output\_name) | The name of the event bus. |
<!-- END_TF_DOCS -->
