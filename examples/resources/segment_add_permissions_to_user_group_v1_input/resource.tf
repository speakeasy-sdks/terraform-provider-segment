resource "segment_add_permissions_to_user_group_v1_input" "my_addpermissionstousergroupv1input" {
  permissions = [
    {
      resources = [
        {
          id = "89bd9d8d-69a6-474e-8f46-7cc8796ed151"
          labels = [
            {
              description = "...my_description..."
              key         = "...my_key..."
              value       = "...my_value..."
            },
          ]
          type = "WAREHOUSE"
        },
      ]
      role_id = "...my_role_id..."
    },
  ]
  user_group_id = "...my_user_group_id..."
}