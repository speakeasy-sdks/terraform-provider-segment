resource "segment_create_destination_v1_input" "my_createdestinationv1input" {
  enabled     = true
  metadata_id = "...my_metadata_id..."
  name        = "Brandon Auer"
  settings = {
    "sed"  = "{ \"see\": \"documentation\" }"
    "iste" = "{ \"see\": \"documentation\" }"
  }
  source_id = "...my_source_id..."
}