resource "segment_create_filter_for_destination_v1_input" "my_createfilterfordestinationv1input" {
  actions = [
    {
      fields = {
        "dolor" = "{ \"see\": \"documentation\" }"
        "natus" = "{ \"see\": \"documentation\" }"
      }
      path    = "...my_path..."
      percent = 38.65
      type    = "SAMPLE"
    },
  ]
  description    = "...my_description..."
  destination_id = "...my_destination_id..."
  enabled        = false
  if             = "...my_if..."
  source_id      = "...my_source_id..."
  title          = "Miss"
}