// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateWarehouseV1OutputLogos - Logo information for this object.
type UpdateWarehouseV1OutputLogos struct {
	// The alternative text for this logo.
	Alt *string `json:"alt,omitempty"`
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark *string `json:"mark,omitempty"`
}

func (o *UpdateWarehouseV1OutputLogos) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *UpdateWarehouseV1OutputLogos) GetDefault() string {
	if o == nil {
		return ""
	}
	return o.Default
}

func (o *UpdateWarehouseV1OutputLogos) GetMark() *string {
	if o == nil {
		return nil
	}
	return o.Mark
}

// UpdateWarehouseV1OutputMetadata - The metadata for the Warehouse.
type UpdateWarehouseV1OutputMetadata struct {
	// A description, in English, of this object.
	Description string `json:"description"`
	// The id of this object.
	ID string `json:"id"`
	// Logo information for this object.
	Logos UpdateWarehouseV1OutputLogos `json:"logos"`
	// The name of this object.
	Name string `json:"name"`
	// The Integration options for this object.
	Options []IntegrationOptionBeta `json:"options"`
	// A human-readable, unique identifier for object.
	Slug string `json:"slug"`
}

func (o *UpdateWarehouseV1OutputMetadata) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *UpdateWarehouseV1OutputMetadata) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateWarehouseV1OutputMetadata) GetLogos() UpdateWarehouseV1OutputLogos {
	if o == nil {
		return UpdateWarehouseV1OutputLogos{}
	}
	return o.Logos
}

func (o *UpdateWarehouseV1OutputMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateWarehouseV1OutputMetadata) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		return []IntegrationOptionBeta{}
	}
	return o.Options
}

func (o *UpdateWarehouseV1OutputMetadata) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

// UpdateWarehouseV1OutputWarehouse - The updated Warehouse.
type UpdateWarehouseV1OutputWarehouse struct {
	// When set to true, this Warehouse receives data.
	Enabled bool `json:"enabled"`
	// The id of the Warehouse.
	ID string `json:"id"`
	// The metadata for the Warehouse.
	Metadata UpdateWarehouseV1OutputMetadata `json:"metadata"`
	// The settings associated with this Warehouse.
	//
	// Common settings are connection-related configuration used to connect to it, for example host, username, and port.
	Settings map[string]interface{} `json:"settings"`
	// The id of the Workspace that owns this Warehouse.
	WorkspaceID string `json:"workspaceId"`
}

func (o *UpdateWarehouseV1OutputWarehouse) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *UpdateWarehouseV1OutputWarehouse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateWarehouseV1OutputWarehouse) GetMetadata() UpdateWarehouseV1OutputMetadata {
	if o == nil {
		return UpdateWarehouseV1OutputMetadata{}
	}
	return o.Metadata
}

func (o *UpdateWarehouseV1OutputWarehouse) GetSettings() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Settings
}

func (o *UpdateWarehouseV1OutputWarehouse) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

// UpdateWarehouseV1Output - Returns the updated Warehouse.
type UpdateWarehouseV1Output struct {
	// The updated Warehouse.
	Warehouse UpdateWarehouseV1OutputWarehouse `json:"warehouse"`
}

func (o *UpdateWarehouseV1Output) GetWarehouse() UpdateWarehouseV1OutputWarehouse {
	if o == nil {
		return UpdateWarehouseV1OutputWarehouse{}
	}
	return o.Warehouse
}
