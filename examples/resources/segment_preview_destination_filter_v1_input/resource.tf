resource "segment_preview_destination_filter_v1_input" "my_previewdestinationfilterv1input" {
  filter = {
    actions = [
      {
        fields = {
          "card"  = "{ \"see\": \"documentation\" }"
          "Haute" = "{ \"see\": \"documentation\" }"
        }
        path    = "...my_path..."
        percent = 16.33
        type    = "DROP_PROPERTIES"
      },
    ]
    if = "...my_if..."
  }
  payload = {
    "Shoes" = "{ \"see\": \"documentation\" }"
    "Auto"  = "{ \"see\": \"documentation\" }"
  }
}