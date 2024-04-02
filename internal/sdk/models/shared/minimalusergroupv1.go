// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MinimalUserGroupV1 - The least amount of information needed to identify a user group.
type MinimalUserGroupV1 struct {
	// The id of the user group.
	ID string `json:"id"`
	// The name of the user group.
	Name string `json:"name"`
}

func (o *MinimalUserGroupV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *MinimalUserGroupV1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}