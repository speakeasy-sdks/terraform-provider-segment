resource "segment_add_permissions_to_user_group_v1_input" "my_addpermissionstousergroupv1input" {
  permissions = [
    {
      resources = [
        {
          id = "94e22b7d-e7cf-4bb2-9dfc-f6e665439410"
          labels = [
            {
              description = "...my_description..."
              key         = "...my_key..."
              value       = "...my_value..."
            },
          ]
          type = "SPACE"
        },
      ]
      role_id = "...my_role_id..."
    },
  ]
  user_group_id = "...my_user_group_id..."
}