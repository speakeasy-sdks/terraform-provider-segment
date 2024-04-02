// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MinimalUserV1 - A user belonging to a Segment Workspace.
type MinimalUserV1 struct {
	// The email address associated with this user.
	Email string `json:"email"`
	// The unique identifier of this user.
	//
	// Config API note: analogous to `name`.
	ID string `json:"id"`
	// The human-readable name of this user.
	Name string `json:"name"`
}

func (o *MinimalUserV1) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *MinimalUserV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *MinimalUserV1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}