// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteFunctionRequest struct {
	FunctionID string `pathParam:"style=simple,explode=false,name=functionId"`
}

func (o *DeleteFunctionRequest) GetFunctionID() string {
	if o == nil {
		return ""
	}
	return o.FunctionID
}

// DeleteFunctionFunctionsResponse200ResponseBody - OK
type DeleteFunctionFunctionsResponse200ResponseBody struct {
	// Delete a single Function.
	Data *shared.DeleteFunctionV1Output `json:"data,omitempty"`
}

func (o *DeleteFunctionFunctionsResponse200ResponseBody) GetData() *shared.DeleteFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteFunctionFunctionsResponseResponseBody - OK
type DeleteFunctionFunctionsResponseResponseBody struct {
	// Delete a single Function.
	Data *shared.DeleteFunctionV1Output `json:"data,omitempty"`
}

func (o *DeleteFunctionFunctionsResponseResponseBody) GetData() *shared.DeleteFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteFunctionFunctionsResponseBody - OK
type DeleteFunctionFunctionsResponseBody struct {
	// Delete a single Function.
	Data *shared.DeleteFunctionV1Output `json:"data,omitempty"`
}

func (o *DeleteFunctionFunctionsResponseBody) GetData() *shared.DeleteFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteFunctionResponseBody - OK
type DeleteFunctionResponseBody struct {
	// Delete a single Function.
	Data *shared.DeleteFunctionV1Output `json:"data,omitempty"`
}

func (o *DeleteFunctionResponseBody) GetData() *shared.DeleteFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type DeleteFunctionResponse struct {
	// OK
	TwoHundredApplicationJSONObject *DeleteFunctionResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *DeleteFunctionFunctionsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *DeleteFunctionFunctionsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *DeleteFunctionFunctionsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteFunctionResponse) GetTwoHundredApplicationJSONObject() *DeleteFunctionResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *DeleteFunctionResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *DeleteFunctionFunctionsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *DeleteFunctionResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *DeleteFunctionFunctionsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *DeleteFunctionResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *DeleteFunctionFunctionsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *DeleteFunctionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteFunctionResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *DeleteFunctionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteFunctionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
