// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
	"net/http"
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

// GetTrackingPlanTrackingPlansResponse200ResponseBody - OK
type GetTrackingPlanTrackingPlansResponse200ResponseBody struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *GetTrackingPlanTrackingPlansResponse200ResponseBody) GetData() *shared.GetTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetTrackingPlanTrackingPlansResponseResponseBody - OK
type GetTrackingPlanTrackingPlansResponseResponseBody struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *GetTrackingPlanTrackingPlansResponseResponseBody) GetData() *shared.GetTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetTrackingPlanTrackingPlansResponseBody - OK
type GetTrackingPlanTrackingPlansResponseBody struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *GetTrackingPlanTrackingPlansResponseBody) GetData() *shared.GetTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetTrackingPlanResponseBody - OK
type GetTrackingPlanResponseBody struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *GetTrackingPlanResponseBody) GetData() *shared.GetTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetTrackingPlanResponse struct {
	// OK
	TwoHundredApplicationJSONObject *GetTrackingPlanResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *GetTrackingPlanTrackingPlansResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *GetTrackingPlanTrackingPlansResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *GetTrackingPlanTrackingPlansResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetTrackingPlanResponse) GetTwoHundredApplicationJSONObject() *GetTrackingPlanResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *GetTrackingPlanTrackingPlansResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *GetTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *GetTrackingPlanTrackingPlansResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *GetTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *GetTrackingPlanTrackingPlansResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
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
