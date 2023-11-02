// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateDestinationV1Input - Updates a single Destination by its id.
type UpdateDestinationV1Input struct {
	// Whether this Destination should receive data.
	Enabled *bool `json:"enabled,omitempty"`
	// Defines the display name of the Destination.
	//
	// Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// An optional object that contains settings for the Destination based on the "required" and "advanced" settings present
	// in the Destination metadata.
	//
	// Config API note: equal to `config`.
	Settings map[string]interface{} `json:"settings,omitempty"`
}

func (o *UpdateDestinationV1Input) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *UpdateDestinationV1Input) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateDestinationV1Input) GetSettings() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Settings
}
