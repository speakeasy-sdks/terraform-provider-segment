// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type UpdateDestinationRequest struct {
	UpdateDestinationV1Input shared.UpdateDestinationV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	DestinationID            string                          `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *UpdateDestinationRequest) GetUpdateDestinationV1Input() shared.UpdateDestinationV1Input {
	if o == nil {
		return shared.UpdateDestinationV1Input{}
	}
	return o.UpdateDestinationV1Input
}

func (o *UpdateDestinationRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

// UpdateDestinationDestinationsResponse200ResponseBody - OK
type UpdateDestinationDestinationsResponse200ResponseBody struct {
	// Returns the updated Destination.
	Data *shared.UpdateDestinationV1Output `json:"data,omitempty"`
}

func (o *UpdateDestinationDestinationsResponse200ResponseBody) GetData() *shared.UpdateDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateDestinationDestinationsResponseResponseBody - OK
type UpdateDestinationDestinationsResponseResponseBody struct {
	// Returns the updated Destination.
	Data *shared.UpdateDestinationV1Output `json:"data,omitempty"`
}

func (o *UpdateDestinationDestinationsResponseResponseBody) GetData() *shared.UpdateDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateDestinationDestinationsResponseBody - OK
type UpdateDestinationDestinationsResponseBody struct {
	// Returns the updated Destination.
	Data *shared.UpdateDestinationV1Output `json:"data,omitempty"`
}

func (o *UpdateDestinationDestinationsResponseBody) GetData() *shared.UpdateDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateDestinationResponseBody - OK
type UpdateDestinationResponseBody struct {
	// Returns the updated Destination.
	Data *shared.UpdateDestinationV1Output `json:"data,omitempty"`
}

func (o *UpdateDestinationResponseBody) GetData() *shared.UpdateDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateDestinationResponse struct {
	// OK
	TwoHundredApplicationJSONObject *UpdateDestinationResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *UpdateDestinationDestinationsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *UpdateDestinationDestinationsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *UpdateDestinationDestinationsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateDestinationResponse) GetTwoHundredApplicationJSONObject() *UpdateDestinationResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *UpdateDestinationResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *UpdateDestinationDestinationsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *UpdateDestinationResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *UpdateDestinationDestinationsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *UpdateDestinationResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *UpdateDestinationDestinationsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *UpdateDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDestinationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *UpdateDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
