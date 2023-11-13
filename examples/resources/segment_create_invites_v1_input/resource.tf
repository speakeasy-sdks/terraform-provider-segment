resource "segment_create_invites_v1_input" "my_createinvitesv1input" {
  invites = [
    {
      email = "Diana.Watsica35@hotmail.com"
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
              id   = "ef356bd2-a2fc-4f13-8b49-24ed7423a910"
              type = "WORKSPACE"
            },
          ]
          role_id = "...my_role_id..."
        },
      ]
    },
  ]
}