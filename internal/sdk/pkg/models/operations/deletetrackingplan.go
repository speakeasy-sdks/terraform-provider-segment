// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type DeleteTrackingPlanRequest struct {
	TrackingPlanID string `pathParam:"style=simple,explode=false,name=trackingPlanId"`
}

func (o *DeleteTrackingPlanRequest) GetTrackingPlanID() string {
	if o == nil {
		return ""
	}
	return o.TrackingPlanID
}

// DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSON - OK
type DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSON struct {
	// Result of a DeleteTrackingPlan call.
	Data *shared.DeleteTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.DeleteTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON - OK
type DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Result of a DeleteTrackingPlan call.
	Data *shared.DeleteTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.DeleteTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSON - OK
type DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSON struct {
	// Result of a DeleteTrackingPlan call.
	Data *shared.DeleteTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSON) GetData() *shared.DeleteTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteTrackingPlan200ApplicationJSON - OK
type DeleteTrackingPlan200ApplicationJSON struct {
	// Result of a DeleteTrackingPlan call.
	Data *shared.DeleteTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *DeleteTrackingPlan200ApplicationJSON) GetData() *shared.DeleteTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type DeleteTrackingPlanResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteTrackingPlan200ApplicationJSONObject *DeleteTrackingPlan200ApplicationJSON
	// OK
	DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSONObject *DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSON
	// OK
	DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject *DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject *DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
}

func (o *DeleteTrackingPlanResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTrackingPlanResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *DeleteTrackingPlanResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTrackingPlanResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteTrackingPlanResponse) GetDeleteTrackingPlan200ApplicationJSONObject() *DeleteTrackingPlan200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteTrackingPlan200ApplicationJSONObject
}

func (o *DeleteTrackingPlanResponse) GetDeleteTrackingPlan200ApplicationVndSegmentV1PlusJSONObject() *DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSONObject
}

func (o *DeleteTrackingPlanResponse) GetDeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject() *DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *DeleteTrackingPlanResponse) GetDeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject() *DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject
}
