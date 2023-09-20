resource "segment_add_labels_to_source_v1_input" "my_addlabelstosourcev1input" {
  labels = [
    {
      description = "...my_description..."
      key         = "...my_key..."
      value       = "...my_value..."
    },
  ]
  source_id = "...my_source_id..."
}