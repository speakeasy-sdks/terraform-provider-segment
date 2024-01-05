---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "segment_create_function_deployment_v1_output Resource - repo"
subcategory: ""
description: |-
  CreateFunctionDeploymentV1Output Resource
---

# segment_create_function_deployment_v1_output (Resource)

CreateFunctionDeploymentV1Output Resource

## Example Usage

```terraform
resource "segment_create_function_deployment_v1_output" "my_createfunctiondeploymentv1output" {
  function_id = "...my_function_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `function_id` (String)

### Read-Only

- `function_deployment` (Attributes) The status of the operation. (see [below for nested schema](#nestedatt--function_deployment))

<a id="nestedatt--function_deployment"></a>
### Nested Schema for `function_deployment`

Read-Only:

- `status` (String) must be one of ["SUCCESS"]

