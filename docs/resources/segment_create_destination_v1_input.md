---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "segment_create_destination_v1_input Resource - repo"
subcategory: ""
description: |-
  CreateDestinationV1Input Resource
---

# segment_create_destination_v1_input (Resource)

CreateDestinationV1Input Resource

## Example Usage

```terraform
resource "segment_create_destination_v1_input" "my_createdestinationv1input" {
  enabled     = true
  metadata_id = "...my_metadata_id..."
  name        = "Dr. Luther Russel"
  settings = {
    "SUV"        = "{ \"see\": \"documentation\" }"
    "synthesize" = "{ \"see\": \"documentation\" }"
  }
  source_id = "...my_source_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata_id` (String) The id of the metadata to link to the new Destination.
- `settings` (Map of String) An object that contains settings for the Destination based on the "required" and "advanced" settings present in the
Destination metadata.

Config API note: equal to `config`.
- `source_id` (String) The id of the Source to connect to this Destination instance.

Config API note: analogous to `parent`.

### Optional

- `enabled` (Boolean) Whether this Destination should receive data.
- `name` (String) Defines the display name of the Destination.

Config API note: equal to `displayName`.

### Read-Only

- `data` (Attributes) Creates a new Destination. (see [below for nested schema](#nestedatt--data))

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `destination` (Attributes) The created Destination. (see [below for nested schema](#nestedatt--data--destination))

<a id="nestedatt--data--destination"></a>
### Nested Schema for `data.destination`

Read-Only:

- `enabled` (Boolean) Whether this instance of a Destination receives data.
- `id` (String) The unique identifier of this instance of a Destination.

Config API note: analogous to `name`.
- `metadata` (Attributes) The metadata of the Destination of which this Destination is an instance of. For example, Google Analytics or Amplitude. (see [below for nested schema](#nestedatt--data--destination--metadata))
- `name` (String) The name of this instance of a Destination.

Config API note: equal to `displayName`.
- `settings` (Map of String) The collection of settings associated with a Destination.

Config API note: equal to `config`.
- `source_id` (String) The id of a Source connected to this instance of a Destination.

Config API note: analogous to `parent`.

<a id="nestedatt--data--destination--metadata"></a>
### Nested Schema for `data.destination.metadata`

Read-Only:

- `actions` (Attributes List) Actions available for the Destination. (see [below for nested schema](#nestedatt--data--destination--metadata--actions))
- `categories` (List of String) A list of categories with which the Destination is associated.
- `components` (Attributes List) A list of components this Destination provides. (see [below for nested schema](#nestedatt--data--destination--metadata--components))
- `contacts` (Attributes List) Contact info for Integration Owners. (see [below for nested schema](#nestedatt--data--destination--metadata--contacts))
- `description` (String) The description of the Destination.
- `id` (String) The id of the Destination metadata.

Config API note: analogous to `name`.
- `logos` (Attributes) The Destination's logos. (see [below for nested schema](#nestedatt--data--destination--metadata--logos))
- `name` (String) The user-friendly name of the Destination.

Config API note: equal to `displayName`.
- `options` (Attributes List) Options configured for the Destination. (see [below for nested schema](#nestedatt--data--destination--metadata--options))
- `partner_owned` (Boolean) Partner Owned flag.
- `presets` (Attributes List) Predefined Destination subscriptions that can optionally be applied when connecting a new instance of the Destination. (see [below for nested schema](#nestedatt--data--destination--metadata--presets))
- `previous_names` (List of String) A list of names previously used by the Destination.
- `region_endpoints` (List of String) The list of regional endpoints for this Destination.
- `slug` (String) The slug used to identify the Destination in the Segment app.
- `status` (String) must be one of ["DEPRECATED", "PRIVATE_BETA", "PRIVATE_BUILDING", "PRIVATE_SUBMITTED", "PUBLIC", "PUBLIC_BETA", "UNAVAILABLE"]
Support status of the Destination.
- `supported_features` (Attributes) Features that this Destination supports.

Config API note: holds `browserUnbundling` fields. (see [below for nested schema](#nestedatt--data--destination--metadata--supported_features))
- `supported_methods` (Attributes) Methods that this Destination supports.

Config API note: equal to `methods`. (see [below for nested schema](#nestedatt--data--destination--metadata--supported_methods))
- `supported_platforms` (Attributes) Platforms from which the Destination receives events.

Config API note: equal to `platforms`. (see [below for nested schema](#nestedatt--data--destination--metadata--supported_platforms))
- `supported_regions` (List of String) A list of supported regions for this Destination.
- `website` (String) A website URL for this Destination.

<a id="nestedatt--data--destination--metadata--actions"></a>
### Nested Schema for `data.destination.metadata.website`

Read-Only:

- `default_trigger` (String) The default value used as the trigger when connecting this action.
- `description` (String) A human-readable description of the action. May include Markdown.
- `fields` (Attributes List) The fields expected in order to perform the action. (see [below for nested schema](#nestedatt--data--destination--metadata--website--fields))
- `hidden` (Boolean) Whether the action should be hidden.
- `id` (String) The primary key of the action.
- `name` (String) A human-readable name for the action.
- `platform` (String) must be one of ["CLOUD", "MOBILE", "WEB"]
The platform on which this action runs.
- `slug` (String) A machine-readable key unique to the action definition.

<a id="nestedatt--data--destination--metadata--website--fields"></a>
### Nested Schema for `data.destination.metadata.website.fields`

Read-Only:

- `allow_null` (Boolean) Whether this field allows null values.
- `choices` (String) Parsed as JSON.
A list of machine-readable value/label pairs to populate a static dropdown.
- `default_value` (String) Parsed as JSON.
A default value that is saved the first time an action is created.
- `description` (String) A human-readable description of this value. You can use Markdown.
- `dynamic` (Boolean) Whether this field should execute a dynamic request to fetch choices to populate a dropdown. When true, `choices` is ignored.
- `field_key` (String) A unique machine-readable key for the field. Should ideally match the expected key in the action\'s API request.
- `id` (String) The primary key of the field.
- `label` (String) A human-readable label for this value.
- `multiple` (Boolean) Whether a user can provide multiples of this field.
- `placeholder` (String) An example value displayed but not saved.
- `required` (Boolean) Whether this field is required.
- `sort_order` (Number) The order this particular field is (used in the UI for displaying the fields in a specified order).
- `type` (String) must be one of ["BOOLEAN", "DATETIME", "HIDDEN", "INTEGER", "NUMBER", "OBJECT", "PASSWORD", "STRING", "TEXT"]
The data type for this value.



<a id="nestedatt--data--destination--metadata--components"></a>
### Nested Schema for `data.destination.metadata.website`

Read-Only:

- `code` (String) Link to the repository hosting the code for this component.
- `owner` (String) must be one of ["PARTNER", "SEGMENT"]
The owner of this component. Either 'SEGMENT' or 'PARTNER'.
- `type` (String) must be one of ["ANDROID", "BROWSER", "IOS", "SERVER"]
The component type.


<a id="nestedatt--data--destination--metadata--contacts"></a>
### Nested Schema for `data.destination.metadata.website`

Read-Only:

- `email` (String) Email of this contact.
- `is_primary` (Boolean) Whether this is a primary contact.
- `name` (String) Name of this contact.
- `role` (String) Role of this contact.


<a id="nestedatt--data--destination--metadata--logos"></a>
### Nested Schema for `data.destination.metadata.website`

Read-Only:

- `alt` (String) The alternative text for this logo.
- `default` (String) The default URL for this logo.
- `mark` (String) The logo mark.


<a id="nestedatt--data--destination--metadata--options"></a>
### Nested Schema for `data.destination.metadata.website`

Read-Only:

- `default_value` (String) Parsed as JSON.
An optional default value for the field.
- `description` (String) An optional short text description of the field.
- `label` (String) An optional label for this field.
- `name` (String) The name identifying this option in the context of a Segment Integration.
- `required` (Boolean) Whether this is a required option when setting up the Integration.
- `type` (String) Defines the type for this option in the schema. Types are most commonly strings, but may also represent other
primitive types, such as booleans, and numbers, as well as complex types, such as objects and arrays.


<a id="nestedatt--data--destination--metadata--presets"></a>
### Nested Schema for `data.destination.metadata.website`

Read-Only:

- `action_id` (String) The unique identifier for the Destination Action to trigger.
- `fields` (Map of String) The default settings for action fields.
- `name` (String) The name of the subscription.
- `trigger` (String) FQL string that describes what events should trigger an action. See https://segment.com/docs/config-api/fql/ for more information regarding Segment's Filter Query Language (FQL).


<a id="nestedatt--data--destination--metadata--supported_features"></a>
### Nested Schema for `data.destination.metadata.website`

Read-Only:

- `browser_unbundling` (Boolean) Whether this Destination supports browser unbundling.
- `browser_unbundling_public` (Boolean) Whether this Destination supports public browser unbundling.
- `cloud_mode_instances` (String) must be one of ["0", "1", "MULTIPLE", "NONE", "SINGLE"]
This Destination's support level for cloud mode instances.
The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
- `device_mode_instances` (String) must be one of ["0", "1", "NONE", "SINGLE"]
This Destination's support level for device mode instances.
Support for multiple device mode instances is currently not planned.
The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
- `replay` (Boolean) Whether this Destination supports replays.


<a id="nestedatt--data--destination--metadata--supported_methods"></a>
### Nested Schema for `data.destination.metadata.website`

Read-Only:

- `alias` (Boolean) Identifies if the Destination supports the `alias` method.
- `group` (Boolean) Identifies if the Destination supports the `group` method.
- `identify` (Boolean) Identifies if the Destination supports the `identify` method.
- `pageview` (Boolean) Identifies if the Destination supports the `pageview` method.
- `track` (Boolean) Identifies if the Destination supports the `track` method.


<a id="nestedatt--data--destination--metadata--supported_platforms"></a>
### Nested Schema for `data.destination.metadata.website`

Read-Only:

- `browser` (Boolean) Whether this Destination supports browser events.
- `mobile` (Boolean) Whether this Destination supports mobile events.
- `server` (Boolean) Whether this Destination supports server events.

