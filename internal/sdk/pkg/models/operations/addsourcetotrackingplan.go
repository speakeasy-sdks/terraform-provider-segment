// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type AddSourceToTrackingPlanRequest struct {
	AddSourceToTrackingPlanV1Input shared.AddSourceToTrackingPlanV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	TrackingPlanID                 string                                `pathParam:"style=simple,explode=false,name=trackingPlanId"`
}

// AddSourceToTrackingPlan200ApplicationVndSegmentV1betaPlusJSON - OK
type AddSourceToTrackingPlan200ApplicationVndSegmentV1betaPlusJSON struct {
	// Connects a Source to a Tracking Plan.
	Data *shared.AddSourceToTrackingPlanV1Output `json:"data,omitempty"`
}

// AddSourceToTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON - OK
type AddSourceToTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Connects a Source to a Tracking Plan.
	Data *shared.AddSourceToTrackingPlanV1Output `json:"data,omitempty"`
}

// AddSourceToTrackingPlan200ApplicationVndSegmentV1PlusJSON - OK
type AddSourceToTrackingPlan200ApplicationVndSegmentV1PlusJSON struct {
	// Connects a Source to a Tracking Plan.
	Data *shared.AddSourceToTrackingPlanV1Output `json:"data,omitempty"`
}

// AddSourceToTrackingPlan200ApplicationJSON - OK
type AddSourceToTrackingPlan200ApplicationJSON struct {
	// Connects a Source to a Tracking Plan.
	Data *shared.AddSourceToTrackingPlanV1Output `json:"data,omitempty"`
}

type AddSourceToTrackingPlanResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	AddSourceToTrackingPlan200ApplicationJSONObject *AddSourceToTrackingPlan200ApplicationJSON
	// OK
	AddSourceToTrackingPlan200ApplicationVndSegmentV1PlusJSONObject *AddSourceToTrackingPlan200ApplicationVndSegmentV1PlusJSON
	// OK
	AddSourceToTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject *AddSourceToTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	AddSourceToTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject *AddSourceToTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
}
