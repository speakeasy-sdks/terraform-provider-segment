resource "segment_preview_destination_filter_v1_input" "my_previewdestinationfilterv1input" {
  filter = {
    actions = [
      {
        fields = {
          "quae"   = "{ \"see\": \"documentation\" }"
          "fugiat" = "{ \"see\": \"documentation\" }"
        }
        path    = "...my_path..."
        percent = 47.71
        type    = "DROP_PROPERTIES"
      },
    ]
    if = "...my_if..."
  }
  payload = {
    "possimus" = "{ \"see\": \"documentation\" }"
    "facilis"  = "{ \"see\": \"documentation\" }"
  }
}