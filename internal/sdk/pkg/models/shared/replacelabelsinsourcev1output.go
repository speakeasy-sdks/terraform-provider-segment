// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ReplaceLabelsInSourceV1Output - Response from replacing a list of labels in a Source.
type ReplaceLabelsInSourceV1Output struct {
	// All labels replaced in the Source.
	Labels []LabelV1 `json:"labels"`
}

func (o *ReplaceLabelsInSourceV1Output) GetLabels() []LabelV1 {
	if o == nil {
		return []LabelV1{}
	}
	return o.Labels
}
