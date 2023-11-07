resource "segment_create_source_v1_input" "my_createsourcev1input" {
  enabled     = true
  metadata_id = "...my_metadata_id..."
  settings = {
    "North"  = "{ \"see\": \"documentation\" }"
    "barren" = "{ \"see\": \"documentation\" }"
  }
  slug = "...my_slug..."
}