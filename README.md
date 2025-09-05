# tf-azurerm-module_primitive-redis_cache_access_policy

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

An access policy that allows configuration of Redis permissions.

## Local Development and Testing

To set yourself up for local development and testing activities, ensure you have the following software available on your PATH:

- make
- git (ensure your user.name and user.email are configured)
- [git-repo](https://gerrit.googlesource.com/git-repo#install)
- [`asdf`](https://asdf-vm.com) or [`mise`](https://mise.jdx.dev/)
- python3 (for pre-commit hooks)

You will also need to authenticate to the Cloud Provider. Terraform will use the default credential resolution mechanism, so ensure you are signed on through the CLI.

Clone this repository to your machine and issue the following command:

```
make configure
```

This will synchronize supporting repositories into this directory and expose additional targets.

To perform linting actions against the Terraform module and Terratests, issue the following command:

```
make lint
```

To provision cloud resources and perform tests against them, issue the following command:

```
make test
```

Note that `make test` causes the creation of some ignored files on your filesystem. This behavior is expected and we want to exclude any state or lockfiles from being pushed to the repository.

These two commands will be utilized in the pipeline and if you cannot run them successfully locally, you are unlikely to see a different result in the pipeline.

For convenience, a target exists that will execute both `make lint` and `make test` for you in sequence. Issue the following command to perform a holistic lint and test:

```
make check
```

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~> 3.117 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [azurerm_redis_cache_access_policy.policy](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/redis_cache_access_policy) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_name"></a> [name](#input\_name) | The name of the Redis Cache Access Policy. Changing this forces a new Redis Cache Access Policy to be created. | `string` | n/a | yes |
| <a name="input_redis_cache_id"></a> [redis\_cache\_id](#input\_redis\_cache\_id) | The ID of the Redis Cache. Changing this forces a new Redis Cache Access Policy to be created. | `string` | n/a | yes |
| <a name="input_permissions"></a> [permissions](#input\_permissions) | Permissions that are going to be assigned to this Redis Cache Access Policy. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | The ID of the Redis Cache Access Policy. |
<!-- END_TF_DOCS -->
