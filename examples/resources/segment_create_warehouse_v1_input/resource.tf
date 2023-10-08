resource "segment_create_warehouse_v1_input" "my_createwarehousev1input" {
  enabled     = true
  metadata_id = "...my_metadata_id..."
  name        = "Marcus Feest"
  settings = {
    "Denar" = "{ \"see\": \"documentation\" }"
    "City"  = "{ \"see\": \"documentation\" }"
  }
}