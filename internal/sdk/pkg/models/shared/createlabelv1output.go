// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateLabelV1OutputLabel - The newly created label.
type CreateLabelV1OutputLabel struct {
	// An optional description of the purpose of this label.
	Description *string `json:"description,omitempty"`
	// The key that represents the name of this label.
	Key string `json:"key"`
	// The value associated with the key of this label.
	Value string `json:"value"`
}

func (o *CreateLabelV1OutputLabel) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateLabelV1OutputLabel) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *CreateLabelV1OutputLabel) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// CreateLabelV1Output - Result of creating a new label in the current Workspace.
type CreateLabelV1Output struct {
	// The newly created label.
	Label CreateLabelV1OutputLabel `json:"label"`
}

func (o *CreateLabelV1Output) GetLabel() CreateLabelV1OutputLabel {
	if o == nil {
		return CreateLabelV1OutputLabel{}
	}
	return o.Label
}
