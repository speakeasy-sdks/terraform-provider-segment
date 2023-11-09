resource "segment_create_workspace_regulation_v1_input" "my_createworkspaceregulationv1input" {
  regulation_type = "DELETE_ONLY"
  subject_ids = [
    "...",
  ]
  subject_type = "USER_ID"
}