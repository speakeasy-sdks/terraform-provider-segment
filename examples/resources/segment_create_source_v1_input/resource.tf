resource "segment_create_source_v1_input" "my_createsourcev1input" {
  enabled     = false
  metadata_id = "...my_metadata_id..."
  settings = {
    "quo"   = "{ \"see\": \"documentation\" }"
    "sequi" = "{ \"see\": \"documentation\" }"
  }
  slug = "...my_slug..."
}