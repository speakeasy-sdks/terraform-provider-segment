resource "segment_create_filter_for_destination_v1_input" "my_createfilterfordestinationv1input" {
  actions = [
    {
      fields = {
        "alarm"      = "{ \"see\": \"documentation\" }"
        "Incredible" = "{ \"see\": \"documentation\" }"
      }
      path    = "...my_path..."
      percent = 98.61
      type    = "ALLOW_PROPERTIES"
    },
  ]
  description    = "...my_description..."
  destination_id = "...my_destination_id..."
  enabled        = false
  if             = "...my_if..."
  source_id      = "...my_source_id..."
  title          = "Mrs."
}