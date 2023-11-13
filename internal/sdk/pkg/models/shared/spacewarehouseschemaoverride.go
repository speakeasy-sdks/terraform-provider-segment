// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SpaceWarehouseSchemaOverride - Overrides the enabled or disabled state of the specified collection and / or properties within the schema.
type SpaceWarehouseSchemaOverride struct {
	// The collection within the Source.
	Collection string `json:"collection"`
	// Represents the overridden enabled state for the listed collection and / or properties.
	Enabled bool `json:"enabled"`
	// A map that contains the properties within the collection to which the Warehouse should sync.
	Property *string `json:"property,omitempty"`
}

func (o *SpaceWarehouseSchemaOverride) GetCollection() string {
	if o == nil {
		return ""
	}
	return o.Collection
}

func (o *SpaceWarehouseSchemaOverride) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *SpaceWarehouseSchemaOverride) GetProperty() *string {
	if o == nil {
		return nil
	}
	return o.Property
}
