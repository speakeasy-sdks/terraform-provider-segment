// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"segment/internal/sdk/pkg/models/operations"
	"segment/internal/sdk/pkg/models/shared"
)

func (r *CreateInvitesV1InputResourceModel) ToCreateSDKType() *shared.CreateInvitesV1Input {
	var invites []shared.InviteV1 = nil
	for _, invitesItem := range r.Invites {
		email := invitesItem.Email.ValueString()
		var permissions []shared.InvitePermissionV1 = nil
		for _, permissionsItem := range invitesItem.Permissions {
			var labels []shared.AllowedLabelBeta = nil
			for _, labelsItem := range permissionsItem.Labels {
				description := new(string)
				if !labelsItem.Description.IsUnknown() && !labelsItem.Description.IsNull() {
					*description = labelsItem.Description.ValueString()
				} else {
					description = nil
				}
				key := labelsItem.Key.ValueString()
				value := labelsItem.Value.ValueString()
				labels = append(labels, shared.AllowedLabelBeta{
					Description: description,
					Key:         key,
					Value:       value,
				})
			}
			var resources []shared.ResourceV1 = nil
			for _, resourcesItem := range permissionsItem.Resources {
				id := resourcesItem.ID.ValueString()
				typeVar := shared.ResourceV1Type(resourcesItem.Type.ValueString())
				resources = append(resources, shared.ResourceV1{
					ID:   id,
					Type: typeVar,
				})
			}
			roleID := permissionsItem.RoleID.ValueString()
			permissions = append(permissions, shared.InvitePermissionV1{
				Labels:    labels,
				Resources: resources,
				RoleID:    roleID,
			})
		}
		invites = append(invites, shared.InviteV1{
			Email:       email,
			Permissions: permissions,
		})
	}
	out := shared.CreateInvitesV1Input{
		Invites: invites,
	}
	return &out
}

func (r *CreateInvitesV1InputResourceModel) RefreshFromCreateResponse(resp *operations.CreateInvitesResponseBody) {
	if resp.Data == nil {
		r.Data = nil
	} else {
		r.Data = &CreateInvitesV1Output{}
		r.Data.Emails = nil
		for _, v := range resp.Data.Emails {
			r.Data.Emails = append(r.Data.Emails, types.StringValue(v))
		}
	}
}
