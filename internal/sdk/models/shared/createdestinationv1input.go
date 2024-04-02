// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateDestinationV1Input - Creates a new Destination.
type CreateDestinationV1Input struct {
	// Whether this Destination should receive data.
	Enabled *bool `json:"enabled,omitempty"`
	// The id of the metadata to link to the new Destination.
	MetadataID string `json:"metadataId"`
	// Defines the display name of the Destination.
	//
	// Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// An object that contains settings for the Destination based on the "required" and "advanced" settings present in the
	// Destination metadata.
	//
	// Config API note: equal to `config`.
	Settings map[string]interface{} `json:"settings"`
	// The id of the Source to connect to this Destination instance.
	//
	// Config API note: analogous to `parent`.
	SourceID string `json:"sourceId"`
}

func (o *CreateDestinationV1Input) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateDestinationV1Input) GetMetadataID() string {
	if o == nil {
		return ""
	}
	return o.MetadataID
}

func (o *CreateDestinationV1Input) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateDestinationV1Input) GetSettings() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Settings
}

func (o *CreateDestinationV1Input) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}