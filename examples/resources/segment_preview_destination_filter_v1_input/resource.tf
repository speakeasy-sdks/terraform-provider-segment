resource "segment_preview_destination_filter_v1_input" "my_previewdestinationfilterv1input" {
  filter = {
    actions = [
      {
        fields = {
          "CSS"   = "{ \"see\": \"documentation\" }"
          "North" = "{ \"see\": \"documentation\" }"
        }
        path    = "...my_path..."
        percent = 38.09
        type    = "ALLOW_PROPERTIES"
      },
    ]
    if = "...my_if..."
  }
  payload = {
    "Supervisor" = "{ \"see\": \"documentation\" }"
    "Northwest"  = "{ \"see\": \"documentation\" }"
  }
}