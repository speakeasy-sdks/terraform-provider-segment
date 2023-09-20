// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type UpdateTrackingPlanRequest struct {
	UpdateTrackingPlanV1Input shared.UpdateTrackingPlanV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	TrackingPlanID            string                           `pathParam:"style=simple,explode=false,name=trackingPlanId"`
}

// UpdateTrackingPlan200ApplicationVndSegmentV1betaPlusJSON - OK
type UpdateTrackingPlan200ApplicationVndSegmentV1betaPlusJSON struct {
	// Result of an UpdateTrackingPlan call.
	Data *shared.UpdateTrackingPlanV1Output `json:"data,omitempty"`
}

// UpdateTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON - OK
type UpdateTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Result of an UpdateTrackingPlan call.
	Data *shared.UpdateTrackingPlanV1Output `json:"data,omitempty"`
}

// UpdateTrackingPlan200ApplicationVndSegmentV1PlusJSON - OK
type UpdateTrackingPlan200ApplicationVndSegmentV1PlusJSON struct {
	// Result of an UpdateTrackingPlan call.
	Data *shared.UpdateTrackingPlanV1Output `json:"data,omitempty"`
}

// UpdateTrackingPlan200ApplicationJSON - OK
type UpdateTrackingPlan200ApplicationJSON struct {
	// Result of an UpdateTrackingPlan call.
	Data *shared.UpdateTrackingPlanV1Output `json:"data,omitempty"`
}

type UpdateTrackingPlanResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	UpdateTrackingPlan200ApplicationJSONObject *UpdateTrackingPlan200ApplicationJSON
	// OK
	UpdateTrackingPlan200ApplicationVndSegmentV1PlusJSONObject *UpdateTrackingPlan200ApplicationVndSegmentV1PlusJSON
	// OK
	UpdateTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject *UpdateTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	UpdateTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject *UpdateTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
}
