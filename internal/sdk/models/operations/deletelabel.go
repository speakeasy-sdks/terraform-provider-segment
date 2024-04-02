// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
	"net/http"
)

type DeleteLabelRequest struct {
	Key   string `pathParam:"style=simple,explode=false,name=key"`
	Value string `pathParam:"style=simple,explode=false,name=value"`
}

func (o *DeleteLabelRequest) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *DeleteLabelRequest) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// DeleteLabelLabelsResponse200ResponseBody - OK
type DeleteLabelLabelsResponse200ResponseBody struct {
	// Returns the status of a label deletion.
	Data *shared.DeleteLabelV1Output `json:"data,omitempty"`
}

func (o *DeleteLabelLabelsResponse200ResponseBody) GetData() *shared.DeleteLabelV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteLabelLabelsResponseResponseBody - OK
type DeleteLabelLabelsResponseResponseBody struct {
	// Returns the status of a label deletion.
	Data *shared.DeleteLabelAlphaOutput `json:"data,omitempty"`
}

func (o *DeleteLabelLabelsResponseResponseBody) GetData() *shared.DeleteLabelAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteLabelLabelsResponseBody - OK
type DeleteLabelLabelsResponseBody struct {
	// Returns the status of a label deletion.
	Data *shared.DeleteLabelV1Output `json:"data,omitempty"`
}

func (o *DeleteLabelLabelsResponseBody) GetData() *shared.DeleteLabelV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteLabelResponseBody - OK
type DeleteLabelResponseBody struct {
	// Returns the status of a label deletion.
	Data *shared.DeleteLabelV1Output `json:"data,omitempty"`
}

func (o *DeleteLabelResponseBody) GetData() *shared.DeleteLabelV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type DeleteLabelResponse struct {
	// OK
	TwoHundredApplicationJSONObject *DeleteLabelResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *DeleteLabelLabelsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *DeleteLabelLabelsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *DeleteLabelLabelsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteLabelResponse) GetTwoHundredApplicationJSONObject() *DeleteLabelResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *DeleteLabelResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *DeleteLabelLabelsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *DeleteLabelResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *DeleteLabelLabelsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *DeleteLabelResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *DeleteLabelLabelsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *DeleteLabelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLabelResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *DeleteLabelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLabelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}