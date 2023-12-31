// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListLabelsAlphaOutput - Returns all available labels for the current Workspace.
type ListLabelsAlphaOutput struct {
	// All labels associated with the current Workspace.
	Labels []LabelAlpha `json:"labels"`
}

func (o *ListLabelsAlphaOutput) GetLabels() []LabelAlpha {
	if o == nil {
		return []LabelAlpha{}
	}
	return o.Labels
}
