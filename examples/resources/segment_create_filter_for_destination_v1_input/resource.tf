resource "segment_create_filter_for_destination_v1_input" "my_createfilterfordestinationv1input" {
  actions = [
    {
      fields = {
        "est"  = "{ \"see\": \"documentation\" }"
        "odit" = "{ \"see\": \"documentation\" }"
      }
      path    = "...my_path..."
      percent = 84.91
      type    = "DROP_PROPERTIES"
    },
  ]
  description    = "...my_description..."
  destination_id = "...my_destination_id..."
  enabled        = true
  if             = "...my_if..."
  source_id      = "...my_source_id..."
  title          = "Dr."
}