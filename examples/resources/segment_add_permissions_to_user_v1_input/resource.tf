resource "segment_add_permissions_to_user_v1_input" "my_addpermissionstouserv1input" {
  permissions = [
    {
      resources = [
        {
          id = "cd6da11d-c06a-4bf6-8342-52089d9c8a56"
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
  user_id = "...my_user_id..."
}