resource "segment_create_invites_v1_input" "my_createinvitesv1input" {
  invites = [
    {
      email = "Madaline21@hotmail.com"
      permissions = [
        {
          labels = [
            {
              description = "...my_description..."
              key         = "...my_key..."
              value       = "...my_value..."
            },
          ]
          resources = [
            {
              id   = "52c59559-07af-4f1a-ba2f-a9467739251a"
              type = "WAREHOUSE"
            },
          ]
          role_id = "...my_role_id..."
        },
      ]
    },
  ]
}