resource "segment_create_destination_v1_input" "my_createdestinationv1input" {
  enabled     = true
  metadata_id = "...my_metadata_id..."
  name        = "June Shields PhD"
  settings = {
    "eveniet"    = "{ \"see\": \"documentation\" }"
    "temporibus" = "{ \"see\": \"documentation\" }"
  }
  source_id = "...my_source_id..."
}