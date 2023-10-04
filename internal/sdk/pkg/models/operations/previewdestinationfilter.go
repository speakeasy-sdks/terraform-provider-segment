// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

// PreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSON - OK
type PreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSON struct {
	// Preview output from applying the Destination filter.
	// Segment modifies or nullifies payloads depending on the provided filter actions.
	Data *shared.PreviewDestinationFilterV1Output `json:"data,omitempty"`
}

// PreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSON - OK
type PreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Preview output from applying the Destination filter.
	// Segment modifies or nullifies payloads depending on the provided filter actions.
	Data *shared.PreviewDestinationFilterV1Output `json:"data,omitempty"`
}

// PreviewDestinationFilter200ApplicationVndSegmentV1PlusJSON - OK
type PreviewDestinationFilter200ApplicationVndSegmentV1PlusJSON struct {
	// Preview output from applying the Destination filter.
	// Segment modifies or nullifies payloads depending on the provided filter actions.
	Data *shared.PreviewDestinationFilterV1Output `json:"data,omitempty"`
}

// PreviewDestinationFilter200ApplicationJSON - OK
type PreviewDestinationFilter200ApplicationJSON struct {
	// Preview output from applying the Destination filter.
	// Segment modifies or nullifies payloads depending on the provided filter actions.
	Data *shared.PreviewDestinationFilterV1Output `json:"data,omitempty"`
}

type PreviewDestinationFilterResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	PreviewDestinationFilter200ApplicationJSONObject *PreviewDestinationFilter200ApplicationJSON
	// OK
	PreviewDestinationFilter200ApplicationVndSegmentV1PlusJSONObject *PreviewDestinationFilter200ApplicationVndSegmentV1PlusJSON
	// OK
	PreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSONObject *PreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	PreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSONObject *PreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSON
}
