// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TransformationBeta - Represents a Transformation.
type TransformationBeta struct {
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
