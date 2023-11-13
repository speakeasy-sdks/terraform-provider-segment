resource "segment_create_warehouse_v1_input" "my_createwarehousev1input" {
  enabled     = true
  metadata_id = "...my_metadata_id..."
  name        = "Margarita Hintz Sr."
  settings = {
    "sievert"     = "{ \"see\": \"documentation\" }"
    "incremental" = "{ \"see\": \"documentation\" }"
  }
}