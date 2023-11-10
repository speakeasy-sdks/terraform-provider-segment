// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AudienceSummary - Defines an Audience.
type AudienceSummary struct {
	// Date the audience was created.
	CreatedAt string `json:"createdAt"`
	// User id who created the audience.
	CreatedBy string `json:"createdBy"`
	// Description of the audience.
	Description string `json:"description"`
	// Enabled/disabled status for the audience.
	Enabled bool `json:"enabled"`
	// Audience id.
	ID string `json:"id"`
	// Key for the audience.
	Key string `json:"key"`
	// Name of the audience.
	Name string `json:"name"`
	// Space id for the audience.
	SpaceID string `json:"spaceId"`
	// Date the audience was last updated.
	UpdatedAt string `json:"updatedAt"`
	// User id who last updated the audience.
	UpdatedBy string `json:"updatedBy"`
}

func (o *AudienceSummary) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *AudienceSummary) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *AudienceSummary) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *AudienceSummary) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *AudienceSummary) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AudienceSummary) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *AudienceSummary) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AudienceSummary) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

func (o *AudienceSummary) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *AudienceSummary) GetUpdatedBy() string {
	if o == nil {
		return ""
	}
	return o.UpdatedBy
}
