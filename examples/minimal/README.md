<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~> 3.117 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | terraform.registry.launch.nttdata.com/module_library/resource_name/launch | ~> 2.0 |
| <a name="module_resource_group"></a> [resource\_group](#module\_resource\_group) | terraform.registry.launch.nttdata.com/module_primitive/resource_group/azurerm | ~> 1.0 |
| <a name="module_redis_cache"></a> [redis\_cache](#module\_redis\_cache) | terraform.registry.launch.nttdata.com/module_primitive/redis_cache/azurerm | ~> 1.1 |
| <a name="module_access_policy"></a> [access\_policy](#module\_access\_policy) | ../.. | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | A map of key to resource\_name that will be used by tf-launch-module\_library-resource\_name to generate resource names | <pre>map(object({<br/>    name       = string<br/>    max_length = optional(number, 60)<br/>    region     = optional(string, "eastus2")<br/>  }))</pre> | <pre>{<br/>  "access_policy": {<br/>    "name": "accesspolicy"<br/>  },<br/>  "redis_cache": {<br/>    "name": "cache"<br/>  },<br/>  "rg": {<br/>    "name": "rg"<br/>  }<br/>}</pre> | no |
| <a name="input_resource_names_strategy"></a> [resource\_names\_strategy](#input\_resource\_names\_strategy) | Strategy to use for generating resource names, taken from the outputs of the naming module, e.g. 'standard', 'minimal\_random\_suffix', 'dns\_compliant\_standard', etc. This example forces the 'minimal\_random\_suffix\_without\_any\_separators' strategy for the ACR and scope map, as those names require alphanumeric characters only. | `string` | `"minimal_random_suffix"` | no |
| <a name="input_logical_product_family"></a> [logical\_product\_family](#input\_logical\_product\_family) | (Required) Name of the product family for which the resource is created.<br/>    Example: org\_name, department\_name. | `string` | `"launch"` | no |
| <a name="input_logical_product_service"></a> [logical\_product\_service](#input\_logical\_product\_service) | (Required) Name of the product service for which the resource is created.<br/>    For example, backend, frontend, middleware etc. | `string` | `"example"` | no |
| <a name="input_class_env"></a> [class\_env](#input\_class\_env) | (Required) Environment where resource is going to be deployed. For example: dev, qa, uat | `string` | `"sandbox"` | no |
| <a name="input_instance_env"></a> [instance\_env](#input\_instance\_env) | Number that represents the instance of the environment. | `number` | `0` | no |
| <a name="input_instance_resource"></a> [instance\_resource](#input\_instance\_resource) | Number that represents the instance of the resource. | `number` | `0` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Tags to apply to all resources. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_resource_group_id"></a> [resource\_group\_id](#output\_resource\_group\_id) | n/a |
| <a name="output_resource_group_name"></a> [resource\_group\_name](#output\_resource\_group\_name) | n/a |
| <a name="output_redis_cache_id"></a> [redis\_cache\_id](#output\_redis\_cache\_id) | n/a |
| <a name="output_redis_cache_name"></a> [redis\_cache\_name](#output\_redis\_cache\_name) | n/a |
| <a name="output_access_policy_id"></a> [access\_policy\_id](#output\_access\_policy\_id) | n/a |
| <a name="output_access_policy_name"></a> [access\_policy\_name](#output\_access\_policy\_name) | n/a |
<!-- END_TF_DOCS -->
