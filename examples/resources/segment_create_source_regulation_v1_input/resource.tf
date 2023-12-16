resource "segment_create_source_regulation_v1_input" "my_createsourceregulationv1input" {
  regulation_type = "SUPPRESS_ONLY"
  source_id       = "...my_source_id..."
  subject_ids = [
    "...",
  ]
  subject_type = "USER_ID"
}