resource "segment_create_function_v1_input" "my_createfunctionv1input" {
  code          = "...my_code..."
  description   = "...my_description..."
  display_name  = "...my_display_name..."
  logo_url      = "...my_logo_url..."
  resource_type = "DESTINATION"
  settings = [
    {
      description = "...my_description..."
      label       = "...my_label..."
      name        = "Shane Herzog"
      required    = false
      sensitive   = true
      type        = "STRING"
    },
  ]
}