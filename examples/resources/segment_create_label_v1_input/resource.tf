resource "segment_create_label_v1_input" "my_createlabelv1input" {
  label = {
    description = "...my_description..."
    key         = "...my_key..."
    value       = "...my_value..."
  }
}