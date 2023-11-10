resource "segment_add_permissions_to_user_group_v1_input" "my_addpermissionstousergroupv1input" {
  permissions = [
    {
      resources = [
        {
          id = "da097c3d-14ba-4809-84bf-04f02d149748"
          labels = [
            {
              description = "...my_description..."
              key         = "...my_key..."
              value       = "...my_value..."
            },
          ]
          type = "SOURCE"
        },
      ]
      role_id = "...my_role_id..."
    },
  ]
  user_group_id = "...my_user_group_id..."
}