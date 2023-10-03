// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type UpdateRulesInTrackingPlanRequest struct {
	UpdateRulesInTrackingPlanV1Input shared.UpdateRulesInTrackingPlanV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	TrackingPlanID                   string                                  `pathParam:"style=simple,explode=false,name=trackingPlanId"`
}

// UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON - OK
type UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

// UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON - OK
type UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

// UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON - OK
type UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

// UpdateRulesInTrackingPlan200ApplicationJSON - OK
type UpdateRulesInTrackingPlan200ApplicationJSON struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

type UpdateRulesInTrackingPlanResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UpdateRulesInTrackingPlan200ApplicationJSONObject *UpdateRulesInTrackingPlan200ApplicationJSON
	// OK
	UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSONObject *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON
	// OK
	UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
}
