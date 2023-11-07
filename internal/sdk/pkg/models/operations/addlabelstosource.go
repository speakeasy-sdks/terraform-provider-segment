// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type AddLabelsToSourceRequest struct {
	AddLabelsToSourceV1Input shared.AddLabelsToSourceV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	SourceID                 string                          `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *AddLabelsToSourceRequest) GetAddLabelsToSourceV1Input() shared.AddLabelsToSourceV1Input {
	if o == nil {
		return shared.AddLabelsToSourceV1Input{}
	}
	return o.AddLabelsToSourceV1Input
}

func (o *AddLabelsToSourceRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// AddLabelsToSourceSourcesResponse200ResponseBody - OK
type AddLabelsToSourceSourcesResponse200ResponseBody struct {
	// Applies an existing label to an existing Source.
	Data *shared.AddLabelsToSourceV1Output `json:"data,omitempty"`
}

func (o *AddLabelsToSourceSourcesResponse200ResponseBody) GetData() *shared.AddLabelsToSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// AddLabelsToSourceSourcesResponseResponseBody - OK
type AddLabelsToSourceSourcesResponseResponseBody struct {
	// Applies an existing label to an existing Source.
	Data *shared.AddLabelsToSourceAlphaOutput `json:"data,omitempty"`
}

func (o *AddLabelsToSourceSourcesResponseResponseBody) GetData() *shared.AddLabelsToSourceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

// AddLabelsToSourceSourcesResponseBody - OK
type AddLabelsToSourceSourcesResponseBody struct {
	// Applies an existing label to an existing Source.
	Data *shared.AddLabelsToSourceV1Output `json:"data,omitempty"`
}

func (o *AddLabelsToSourceSourcesResponseBody) GetData() *shared.AddLabelsToSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// AddLabelsToSourceResponseBody - OK
type AddLabelsToSourceResponseBody struct {
	// Applies an existing label to an existing Source.
	Data *shared.AddLabelsToSourceV1Output `json:"data,omitempty"`
}

func (o *AddLabelsToSourceResponseBody) GetData() *shared.AddLabelsToSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type AddLabelsToSourceResponse struct {
	// OK
	TwoHundredApplicationJSONObject *AddLabelsToSourceResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *AddLabelsToSourceSourcesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *AddLabelsToSourceSourcesResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *AddLabelsToSourceSourcesResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AddLabelsToSourceResponse) GetTwoHundredApplicationJSONObject() *AddLabelsToSourceResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *AddLabelsToSourceResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *AddLabelsToSourceSourcesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *AddLabelsToSourceResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *AddLabelsToSourceSourcesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *AddLabelsToSourceResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *AddLabelsToSourceSourcesResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *AddLabelsToSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddLabelsToSourceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *AddLabelsToSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddLabelsToSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
