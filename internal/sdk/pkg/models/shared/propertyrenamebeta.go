// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PropertyRenameBeta struct {
	// The new name to rename the property.
	NewName string `json:"newName"`
	// The old name of the property.
	OldName string `json:"oldName"`
}

func (o *PropertyRenameBeta) GetNewName() string {
	if o == nil {
		return ""
	}
	return o.NewName
}

func (o *PropertyRenameBeta) GetOldName() string {
	if o == nil {
		return ""
	}
	return o.OldName
}
