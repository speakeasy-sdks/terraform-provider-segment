resource "segment_create_destination_v1_input" "my_createdestinationv1input" {
  enabled     = true
  metadata_id = "...my_metadata_id..."
  name        = "June Shields PhD"
  settings = {
    "wherever" = "{ \"see\": \"documentation\" }"
    "Direct"   = "{ \"see\": \"documentation\" }"
  }
  source_id = "...my_source_id..."
}