resource "segment_create_transformation_v1_input" "my_createtransformationv1input" {
  destination_metadata_id = "...my_destination_metadata_id..."
  enabled                 = false
  fql_defined_properties = [
    {
      fql           = "...my_fql..."
      property_name = "...my_property_name..."
    },
  ]
  if             = "...my_if..."
  name           = "Tony Dare"
  new_event_name = "...my_new_event_name..."
  property_renames = [
    {
      new_name = "...my_new_name..."
      old_name = "...my_old_name..."
    },
  ]
  property_value_transformations = [
    {
      property_paths = [
        "...",
      ]
      property_value = "...my_property_value..."
    },
  ]
  source_id = "...my_source_id..."
}