// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListRegulationsFromSourceV1Output - Output of all Source-scoped regulations.
type ListRegulationsFromSourceV1Output struct {
	// List of Workspace-scoped regulations with statuses.
	Regulations []RegulationListEntryV1 `json:"regulations"`
}

func (o *ListRegulationsFromSourceV1Output) GetRegulations() []RegulationListEntryV1 {
	if o == nil {
		return []RegulationListEntryV1{}
	}
	return o.Regulations
}
