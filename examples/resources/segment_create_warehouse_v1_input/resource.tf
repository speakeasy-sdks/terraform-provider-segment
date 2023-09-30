resource "segment_create_warehouse_v1_input" "my_createwarehousev1input" {
  enabled     = false
  metadata_id = "...my_metadata_id..."
  name        = "Wendy Hansen"
  settings = {
    "inventore" = "{ \"see\": \"documentation\" }"
    "dolorem"   = "{ \"see\": \"documentation\" }"
  }
}