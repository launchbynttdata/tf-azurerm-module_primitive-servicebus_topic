# tf-azurerm-module_primitive-servicebus_topic

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This module is managed via the Launch Terraform skeleton.

## Usage

See [examples/complete](examples/complete) for a full working example.

## Module Development

### Pre-Requisites

The following commands should be available on your system:

- `asdf` or `mise`
- `make`
- `python3` (for pre-commit)

Additionally, your `git` user and email must be configured. Run the `make configure` command from the root of the repository to ensure that you meet these requirements.

### Pre-Commit hooks

The [.pre-commit-config.yaml](.pre-commit-config.yaml) file defines `pre-commit` hooks for Terraform formatting, validation, documentation generation, and detect-secrets. Hooks are installed when you run `make configure`. Go linting runs via `make lint` in local development and CI, not via pre-commit.

### Terratest examples

Post-deploy tests in `tests/post_deploy_functional/` and `tests/post_deploy_functional_readonly/` target `examples/complete` via an explicit folder constant in each `main_test.go`. Adding another example requires a new test entry point or updating that constant; it is not picked up automatically.

### Local Validation

You should validate the changes you make to any module locally, prior to pushing your changes in a branch to GitHub.

1. Ensure that you have run `make configure` successfully.
2. Ensure you are signed into the appropriate cloud provider (e.g. Azure) for the module under test in your current console session.
3. Run the Terraform and Golang linters:

```
make lint
```

4. Once linters pass, run integration tests (apply, test, destroy):

```
make test
```

The pre-commit validations, as well as the `make lint` and `make test` targets, are performed in CI. Running them locally before opening a PR helps ensure a smooth review.

### Review & Merge Process

Open a Pull Request to the default (`main`) branch. The PR title must follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/#specification) format to merge and to drive semantic versioning.

Ensure CI workflows pass, address review feedback, and obtain approvals required by `CODEOWNERS`.

### Automatic Updates

Shared configuration and workflow files are largely managed through [launch-terraform-skeleton](https://github.com/launchbynttdata/launch-terraform-skeleton). Avoid one-off edits to copied skeleton files in this repository unless necessary (for example `.gitignore` entries for generated artifacts). Use `copier check-update` / `copier update` when refreshing from the skeleton.


<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~> 3.77 |

## Modules

No modules.

## Resources

| Name | Type |
|------|---------|
| [azurerm_servicebus_topic.topic](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/servicebus_topic) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_auto_delete_on_idle"></a> [auto\_delete\_on\_idle](#input\_auto\_delete\_on\_idle) | The idle interval after which the topic is automatically deleted. | `string` | `"P10675199DT2H48M5.4775807S"` | no |
| <a name="input_default_message_ttl"></a> [default\_message\_ttl](#input\_default\_message\_ttl) | The TTL of messages sent to this topic if no TTL value is set on the message itself. | `string` | `"P10675199DT2H48M5.4775807S"` | no |
| <a name="input_duplicate_detection_history_time_window"></a> [duplicate\_detection\_history\_time\_window](#input\_duplicate\_detection\_history\_time\_window) | The duration of the duplicate detection history. | `string` | `"PT10M"` | no |
| <a name="input_max_size_in_megabytes"></a> [max\_size\_in\_megabytes](#input\_max\_size\_in\_megabytes) | The maximum size of the topic in megabytes. | `number` | `1024` | no |
| <a name="input_name"></a> [name](#input\_name) | The name of the Service Bus Topic. | `string` | n/a | yes |
| <a name="input_namespace_id"></a> [namespace\_id](#input\_namespace\_id) | The ID of the Service Bus Namespace. | `string` | n/a | yes |
| <a name="input_partitioning_enabled"></a> [partitioning\_enabled](#input\_partitioning\_enabled) | Boolean flag which controls whether to enable partitioning on the Service Bus Topic. | `bool` | `true` | no |
| <a name="input_status"></a> [status](#input\_status) | The status of the Service Bus Topic. | `string` | `"Active"` | no |
| <a name="input_support_ordering"></a> [support\_ordering](#input\_support\_ordering) | Boolean flag which controls whether the topic supports ordering. | `bool` | `false` | no |

## Outputs

| Name | Description |
|------|---------|
| <a name="output_id"></a> [id](#output\_id) | The ID of the Service Bus topic |
| <a name="output_name"></a> [name](#output\_name) | The name of the Service Bus topic |
<!-- END_TF_DOCS -->