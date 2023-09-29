resource "segment_create_invites_v1_input" "my_createinvitesv1input" {
  invites = [
    {
      email = "Florian.Hilll@yahoo.com"
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
              id   = "98f5b93c-18bd-4388-8e54-e4dac1a06be8"
              type = "WAREHOUSE"
            },
          ]
          role_id = "...my_role_id..."
        },
      ]
    },
  ]
}