// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PropertyRenameV1 struct {
	// The new name to rename the property.
	NewName string `json:"newName"`
	// The old name of the property.
	OldName string `json:"oldName"`
}

func (o *PropertyRenameV1) GetNewName() string {
	if o == nil {
		return ""
	}
	return o.NewName
}

func (o *PropertyRenameV1) GetOldName() string {
	if o == nil {
		return ""
	}
	return o.OldName
}
