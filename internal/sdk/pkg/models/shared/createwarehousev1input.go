// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateWarehouseV1Input - Create a new Warehouse based on a set of parameters.
type CreateWarehouseV1Input struct {
	// Enable to allow this Warehouse to receive data. Defaults to true.
	Enabled *bool `json:"enabled,omitempty"`
	// The Warehouse metadata to use.
	MetadataID string `json:"metadataId"`
	// An optional human-readable name for this Warehouse.
	Name *string `json:"name,omitempty"`
	// A key-value object that contains instance-specific settings for a Warehouse.
	//
	// Different kinds of Warehouses require different settings. The required and optional settings
	// for a Warehouse are described in the `options` object of the associated Warehouse metadata.
	//
	// You can find the full list of Warehouse metadata and related settings information in the
	// `/catalog/warehouses` endpoint.
	Settings map[string]interface{} `json:"settings"`
}

func (o *CreateWarehouseV1Input) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateWarehouseV1Input) GetMetadataID() string {
	if o == nil {
		return ""
	}
	return o.MetadataID
}

func (o *CreateWarehouseV1Input) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateWarehouseV1Input) GetSettings() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Settings
}
