// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetWarehouseMetadataV1OutputWarehouseMetadataV1LogosBeta - Represents a logo.
type GetWarehouseMetadataV1OutputWarehouseMetadataV1LogosBeta struct {
	// The alternative text for this logo.
	Alt *string `json:"alt,omitempty"`
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark *string `json:"mark,omitempty"`
}

func (o *GetWarehouseMetadataV1OutputWarehouseMetadataV1LogosBeta) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *GetWarehouseMetadataV1OutputWarehouseMetadataV1LogosBeta) GetDefault() string {
	if o == nil {
		return ""
	}
	return o.Default
}

func (o *GetWarehouseMetadataV1OutputWarehouseMetadataV1LogosBeta) GetMark() *string {
	if o == nil {
		return nil
	}
	return o.Mark
}

// GetWarehouseMetadataV1OutputWarehouseMetadataV1 - The metadata for an instance of Segment’s data Warehouse product.
type GetWarehouseMetadataV1OutputWarehouseMetadataV1 struct {
	// A description, in English, of this object.
	Description string `json:"description"`
	// The id of this object.
	ID string `json:"id"`
	// Logo information for this object.
	Logos GetWarehouseMetadataV1OutputWarehouseMetadataV1LogosBeta `json:"logos"`
	// The name of this object.
	Name string `json:"name"`
	// The Integration options for this object.
	Options []IntegrationOptionBeta `json:"options"`
	// A human-readable, unique identifier for object.
	Slug string `json:"slug"`
}

func (o *GetWarehouseMetadataV1OutputWarehouseMetadataV1) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *GetWarehouseMetadataV1OutputWarehouseMetadataV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetWarehouseMetadataV1OutputWarehouseMetadataV1) GetLogos() GetWarehouseMetadataV1OutputWarehouseMetadataV1LogosBeta {
	if o == nil {
		return GetWarehouseMetadataV1OutputWarehouseMetadataV1LogosBeta{}
	}
	return o.Logos
}

func (o *GetWarehouseMetadataV1OutputWarehouseMetadataV1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetWarehouseMetadataV1OutputWarehouseMetadataV1) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		return []IntegrationOptionBeta{}
	}
	return o.Options
}

func (o *GetWarehouseMetadataV1OutputWarehouseMetadataV1) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

// GetWarehouseMetadataV1Output - Returns a Warehouse catalog item looked up by id.
type GetWarehouseMetadataV1Output struct {
	// The Warehouse catalog item.
	WarehouseMetadata GetWarehouseMetadataV1OutputWarehouseMetadataV1 `json:"warehouseMetadata"`
}

func (o *GetWarehouseMetadataV1Output) GetWarehouseMetadata() GetWarehouseMetadataV1OutputWarehouseMetadataV1 {
	if o == nil {
		return GetWarehouseMetadataV1OutputWarehouseMetadataV1{}
	}
	return o.WarehouseMetadata
}
