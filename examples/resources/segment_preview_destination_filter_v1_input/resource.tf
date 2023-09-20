resource "segment_preview_destination_filter_v1_input" "my_previewdestinationfilterv1input" {
  filter = {
    actions = [
      {
        fields = {
          "accusamus" = "{ \"see\": \"documentation\" }"
          "commodi"   = "{ \"see\": \"documentation\" }"
        }
        path    = "...my_path..."
        percent = 91.83
        type    = "ALLOW_PROPERTIES"
      },
    ]
    if = "...my_if..."
  }
  payload = {
    "ipsum"  = "{ \"see\": \"documentation\" }"
    "quidem" = "{ \"see\": \"documentation\" }"
  }
}