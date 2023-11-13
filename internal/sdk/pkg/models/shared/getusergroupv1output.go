// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetUserGroupV1OutputUserGroupV1 - A set of users with a set of shared permissions.
type GetUserGroupV1OutputUserGroupV1 struct {
	// The id of the user group.
	ID string `json:"id"`
	// The number of members in the user group.
	MemberCount float64 `json:"memberCount"`
	// The name of the user group.
	Name string `json:"name"`
	// The permissions associated with the user group.
	Permissions []PermissionV1 `json:"permissions,omitempty"`
}

func (o *GetUserGroupV1OutputUserGroupV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetUserGroupV1OutputUserGroupV1) GetMemberCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.MemberCount
}

func (o *GetUserGroupV1OutputUserGroupV1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetUserGroupV1OutputUserGroupV1) GetPermissions() []PermissionV1 {
	if o == nil {
		return nil
	}
	return o.Permissions
}

// GetUserGroupV1Output - Returns a user group with the given id.
type GetUserGroupV1Output struct {
	// The user group returned from the query.
	UserGroup GetUserGroupV1OutputUserGroupV1 `json:"userGroup"`
}

func (o *GetUserGroupV1Output) GetUserGroup() GetUserGroupV1OutputUserGroupV1 {
	if o == nil {
		return GetUserGroupV1OutputUserGroupV1{}
	}
	return o.UserGroup
}
