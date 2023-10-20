// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type GetTrackingPlanRequest struct {
	TrackingPlanID string `pathParam:"style=simple,explode=false,name=trackingPlanId"`
}

func (o *GetTrackingPlanRequest) GetTrackingPlanID() string {
	if o == nil {
		return ""
	}
	return o.TrackingPlanID
}

// GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSON - OK
type GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSON struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.GetTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.GetTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetTrackingPlan200ApplicationVndSegmentV1PlusJSON - OK
type GetTrackingPlan200ApplicationVndSegmentV1PlusJSON struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *GetTrackingPlan200ApplicationVndSegmentV1PlusJSON) GetData() *shared.GetTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetTrackingPlan200ApplicationJSON - OK
type GetTrackingPlan200ApplicationJSON struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *GetTrackingPlan200ApplicationJSON) GetData() *shared.GetTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetTrackingPlanResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GetTrackingPlan200ApplicationJSONObject *GetTrackingPlan200ApplicationJSON
	// OK
	GetTrackingPlan200ApplicationVndSegmentV1PlusJSONObject *GetTrackingPlan200ApplicationVndSegmentV1PlusJSON
	// OK
	GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject *GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject *GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
}

func (o *GetTrackingPlanResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTrackingPlanResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetTrackingPlanResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTrackingPlanResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTrackingPlanResponse) GetGetTrackingPlan200ApplicationJSONObject() *GetTrackingPlan200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetTrackingPlan200ApplicationJSONObject
}

func (o *GetTrackingPlanResponse) GetGetTrackingPlan200ApplicationVndSegmentV1PlusJSONObject() *GetTrackingPlan200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.GetTrackingPlan200ApplicationVndSegmentV1PlusJSONObject
}

func (o *GetTrackingPlanResponse) GetGetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject() *GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *GetTrackingPlanResponse) GetGetTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject() *GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject
}
