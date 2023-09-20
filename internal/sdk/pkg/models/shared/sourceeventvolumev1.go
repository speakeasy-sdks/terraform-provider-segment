// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SourceEventVolumeV1EventSourceV1 - Source represents a Segment Source.
type SourceEventVolumeV1EventSourceV1 struct {
	// The id of the Source where the events came from.
	ID string `json:"id"`
	// The name of the Source, if applicable.
	Name *string `json:"name,omitempty"`
	// The slug of the Source, if applicable.
	Slug *string `json:"slug,omitempty"`
}

// SourceEventVolumeV1 - SourceEventVolume represents a time series of event volume for a Workspace
// broken down by the dimensions which the customer specifies (optional
// parameters).
type SourceEventVolumeV1 struct {
	// The name of the event, if applicable.
	EventName *string `json:"eventName,omitempty"`
	// The event type, if applicable.
	EventType *string `json:"eventType,omitempty"`
	// A list of the event counts broken down by the requested
	// granularity.
	Series []SourceEventVolumeDatapointV1 `json:"series"`
	// The Source where the events originated.
	Source *SourceEventVolumeV1EventSourceV1 `json:"source,omitempty"`
	// The total count for all events in the requested time frame.
	Total float64 `json:"total"`
}