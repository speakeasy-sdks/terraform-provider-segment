resource "segment_create_warehouse_v1_input" "my_createwarehousev1input" {
  enabled     = false
  metadata_id = "...my_metadata_id..."
  name        = "Mrs. April Wuckert"
  settings = {
    "iusto" = "{ \"see\": \"documentation\" }"
    "dicta" = "{ \"see\": \"documentation\" }"
  }
}