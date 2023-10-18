// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateTransformationBetaOutputTransformationBeta - Represents a Transformation.
type CreateTransformationBetaOutputTransformationBeta struct {
	// The optional Destination metadata associated with the Transformation.
	DestinationMetadataID *string `json:"destinationMetadataId,omitempty"`
	// If the Transformation is enabled.
	Enabled bool `json:"enabled"`
	// The id of the Transformation.
	ID string `json:"id"`
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
	PropertyRenames []PropertyRenameBeta `json:"propertyRenames,omitempty"`
	// Optional array for transforming properties and values collected by your events. Limited to 10 properties.
	PropertyValueTransformations []PropertyValueTransformationBeta `json:"propertyValueTransformations,omitempty"`
	// The Source associated with the Transformation.
	SourceID string `json:"sourceId"`
}

func (o *CreateTransformationBetaOutputTransformationBeta) GetDestinationMetadataID() *string {
	if o == nil {
		return nil
	}
	return o.DestinationMetadataID
}

func (o *CreateTransformationBetaOutputTransformationBeta) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *CreateTransformationBetaOutputTransformationBeta) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateTransformationBetaOutputTransformationBeta) GetIf() string {
	if o == nil {
		return ""
	}
	return o.If
}

func (o *CreateTransformationBetaOutputTransformationBeta) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateTransformationBetaOutputTransformationBeta) GetNewEventName() *string {
	if o == nil {
		return nil
	}
	return o.NewEventName
}

func (o *CreateTransformationBetaOutputTransformationBeta) GetPropertyRenames() []PropertyRenameBeta {
	if o == nil {
		return nil
	}
	return o.PropertyRenames
}

func (o *CreateTransformationBetaOutputTransformationBeta) GetPropertyValueTransformations() []PropertyValueTransformationBeta {
	if o == nil {
		return nil
	}
	return o.PropertyValueTransformations
}

func (o *CreateTransformationBetaOutputTransformationBeta) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// CreateTransformationBetaOutput - The output of a created Transformation.
type CreateTransformationBetaOutput struct {
	// The created Transformation.
	Transformation CreateTransformationBetaOutputTransformationBeta `json:"transformation"`
}

func (o *CreateTransformationBetaOutput) GetTransformation() CreateTransformationBetaOutputTransformationBeta {
	if o == nil {
		return CreateTransformationBetaOutputTransformationBeta{}
	}
	return o.Transformation
}
