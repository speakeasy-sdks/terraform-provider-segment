resource "segment_create_cloud_source_regulation_v1_input" "my_createcloudsourceregulationv1input" {
  collection      = "...my_collection..."
  regulation_type = "SUPPRESS_ONLY"
  source_id       = "...my_source_id..."
  subject_ids = [
    "...",
  ]
  subject_type = "OBJECT_ID"
}