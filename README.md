# Terraform AWS Module: CloudWatch Event Bus (Primitive)

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This Terraform module wraps the [aws_cloudwatch_event_bus](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_bus) resource to create an Amazon EventBridge (CloudWatch Events) event bus. It exposes all documented capabilities of the resource including optional KMS encryption, dead-letter queue configuration, and logging.

EventBridge was formerly known as CloudWatch Events; the functionality is identical.

## Pre-Commit Hooks

The [.pre-commit-config.yaml](.pre-commit-config.yaml) file defines `pre-commit` hooks for Terraform, Go, and common linting tasks.

`commitlint` enforces conventional commit message format. The `detect-secrets-hook` prevents new secrets from being introduced into the baseline. See the [pre-commit documentation](https://pre-commit.com/) for installation and usage.

To enable the commit-msg hook for commitlint:

```
pre-commit install --hook-type commit-msg
```

## Usage

```hcl
module "event_bus" {
  source = "terraform.registry.launch.nttdata.com/module_primitive/cloudwatch_event_bus/aws"
  version = "~> 1.0"

  name               = "my-event-bus"
  description        = "Event bus for application events"
  kms_key_identifier = aws_kms_key.event_bus.arn
  tags               = var.tags
}
```

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.14 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_cloudwatch_event_bus.event_bus](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_bus) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_dead_letter_config"></a> [dead\_letter\_config](#input\_dead\_letter\_config) | Configuration details of the Amazon SQS queue for EventBridge to use as a dead-letter queue (DLQ). | <pre>object({<br/>    arn = string<br/>  })</pre> | `null` | no |
| <a name="input_description"></a> [description](#input\_description) | Event bus description. | `string` | `null` | no |
| <a name="input_event_source_name"></a> [event\_source\_name](#input\_event\_source\_name) | Partner event source that the new event bus will be matched with. Must match name for partner event buses. | `string` | `null` | no |
| <a name="input_kms_key_identifier"></a> [kms\_key\_identifier](#input\_kms\_key\_identifier) | Identifier of the AWS KMS customer managed key for EventBridge to use for encrypting events. Can be the key ARN, KeyId, key alias, or key alias ARN. | `string` | `null` | no |
| <a name="input_name"></a> [name](#input\_name) | Name of the new event bus. The names of custom event buses cannot contain the / character. To create a partner event bus, ensure that the name matches the event\_source\_name. | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | Map of tags to assign to the event bus. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the event bus. |
| <a name="output_id"></a> [id](#output\_id) | The ID of the event bus (same as the name). |
| <a name="output_name"></a> [name](#output\_name) | The name of the event bus. |
<!-- END_TF_DOCS -->

## Testing

1. Run `make configure` to install dependencies.
2. For AWS, set up credentials (e.g., `AWS_PROFILE`, `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY`, or `make env` if configured).
3. Create `examples/complete/provider.tf` with AWS provider configuration (delivered by Makefile when using `make configure`).
4. Run `make check` to execute lint, validate, plan, conftests, terratest, and OPA tests.
