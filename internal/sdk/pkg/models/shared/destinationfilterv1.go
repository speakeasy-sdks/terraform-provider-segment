// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DestinationFilterV1 - Represents a Destination filter.
type DestinationFilterV1 struct {
	// A list of actions this filter performs.
	Actions []DestinationFilterActionV1 `json:"actions"`
	// The timestamp of this filter's creation.
	CreatedAt string `json:"createdAt"`
	// A description for this filter.
	Description *string `json:"description,omitempty"`
	// The id of the Destination associated with this filter.
	DestinationID string `json:"destinationId"`
	// When set to true, this filter is active.
	Enabled bool `json:"enabled"`
	// The unique id of this filter.
	ID string `json:"id"`
	// A condition that defines whether to apply this filter to a payload.
	If string `json:"if"`
	// The id of the Source associated with this filter.
	SourceID string `json:"sourceId"`
	// A title for this filter.
	Title string `json:"title"`
	// The timestamp of this filter's last change.
	UpdatedAt string `json:"updatedAt"`
}

func (o *DestinationFilterV1) GetActions() []DestinationFilterActionV1 {
	if o == nil {
		return []DestinationFilterActionV1{}
	}
	return o.Actions
}

func (o *DestinationFilterV1) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *DestinationFilterV1) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *DestinationFilterV1) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *DestinationFilterV1) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *DestinationFilterV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DestinationFilterV1) GetIf() string {
	if o == nil {
		return ""
	}
	return o.If
}

func (o *DestinationFilterV1) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *DestinationFilterV1) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *DestinationFilterV1) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}
