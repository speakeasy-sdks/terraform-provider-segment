resource "segment_add_permissions_to_user_v1_input" "my_addpermissionstouserv1input" {
  permissions = [
    {
      resources = [
        {
          id = "05dfc2dd-f7cc-478c-a1ba-928fc816742c"
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