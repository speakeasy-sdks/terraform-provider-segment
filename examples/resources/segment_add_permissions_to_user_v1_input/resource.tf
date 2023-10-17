resource "segment_add_permissions_to_user_v1_input" "my_addpermissionstouserv1input" {
  permissions = [
    {
      resources = [
        {
          id = "a872f86d-e977-4402-937d-61a4a0144e46"
          labels = [
            {
              description = "...my_description..."
              key         = "...my_key..."
              value       = "...my_value..."
            },
          ]
          type = "FUNCTION"
        },
      ]
      role_id = "...my_role_id..."
    },
  ]
  user_id = "...my_user_id..."
}