resource "segment_create_workspace_regulation_v1_input" "my_createworkspaceregulationv1input" {
  regulation_type = "DELETE_INTERNAL"
  subject_ids = [
    "...",
  ]
  subject_type = "OBJECT_ID"
}