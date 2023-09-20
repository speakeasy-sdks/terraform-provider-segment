resource "segment_label_v1" "my_labelv1" {
  description = "...my_description..."
  key         = "...my_key..."
  labels = [
    {
      description = "...my_description..."
      key         = "...my_key..."
      value       = "...my_value..."
    },
  ]
  source_id = "...my_source_id..."
  value     = "...my_value..."
}