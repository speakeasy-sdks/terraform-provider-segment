// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateTransformationV1Input - The input to create a Transformation.
type CreateTransformationV1Input struct {
	// The optional Destination metadata id to be associated with the Transformation.
	DestinationMetadataID *string `json:"destinationMetadataId,omitempty"`
	// If the Transformation should be enabled.
	Enabled bool `json:"enabled"`
	// Optional array for defining new properties in [FQL](https://segment.com/docs/config-api/fql/). Currently limited to 1 property.
	FqlDefinedProperties []FQLDefinedPropertyV1 `json:"fqlDefinedProperties,omitempty"`
	// If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.
	//
	// For standard event matchers, use the following:
	//   Track -\> "event='\<eventName\>'"
	//   Identify -\> "type='identify'"
	//   Group -\> "type='group'"
	If string `json:"if"`
	// The name of the Transformation.
	Name string `json:"name"`
	// Optional new event name for renaming events. Works only for 'track' event type.
	NewEventName *string `json:"newEventName,omitempty"`
	// Optional array for renaming properties collected by your events.
	PropertyRenames []PropertyRenameV1 `json:"propertyRenames,omitempty"`
	// Optional array for transforming properties and values collected by your events. Limited to 10 properties.
	PropertyValueTransformations []PropertyValueTransformationV1 `json:"propertyValueTransformations,omitempty"`
	// The Source to be associated with the Transformation.
	SourceID string `json:"sourceId"`
}

func (o *CreateTransformationV1Input) GetDestinationMetadataID() *string {
	if o == nil {
		return nil
	}
	return o.DestinationMetadataID
}

func (o *CreateTransformationV1Input) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *CreateTransformationV1Input) GetFqlDefinedProperties() []FQLDefinedPropertyV1 {
	if o == nil {
		return nil
	}
	return o.FqlDefinedProperties
}

func (o *CreateTransformationV1Input) GetIf() string {
	if o == nil {
		return ""
	}
	return o.If
}

func (o *CreateTransformationV1Input) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateTransformationV1Input) GetNewEventName() *string {
	if o == nil {
		return nil
	}
	return o.NewEventName
}

func (o *CreateTransformationV1Input) GetPropertyRenames() []PropertyRenameV1 {
	if o == nil {
		return nil
	}
	return o.PropertyRenames
}

func (o *CreateTransformationV1Input) GetPropertyValueTransformations() []PropertyValueTransformationV1 {
	if o == nil {
		return nil
	}
	return o.PropertyValueTransformations
}

func (o *CreateTransformationV1Input) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}
