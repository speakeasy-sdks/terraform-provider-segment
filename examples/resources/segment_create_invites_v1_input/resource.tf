resource "segment_create_invites_v1_input" "my_createinvitesv1input" {
  invites = [
    {
      email = "Francesca.Reichert@hotmail.com"
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
              id   = "c18bd388-0e54-4e4d-ac1a-06be8aac3463"
              type = "FUNCTION"
            },
          ]
          role_id = "...my_role_id..."
        },
      ]
    },
  ]
}