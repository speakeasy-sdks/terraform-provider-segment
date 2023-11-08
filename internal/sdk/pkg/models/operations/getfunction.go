// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetFunctionRequest struct {
	FunctionID string `pathParam:"style=simple,explode=false,name=functionId"`
}

func (o *GetFunctionRequest) GetFunctionID() string {
	if o == nil {
		return ""
	}
	return o.FunctionID
}

// GetFunctionFunctionsResponse200ResponseBody - OK
type GetFunctionFunctionsResponse200ResponseBody struct {
	// Gets a single Function.
	Data *shared.GetFunctionV1Output `json:"data,omitempty"`
}

func (o *GetFunctionFunctionsResponse200ResponseBody) GetData() *shared.GetFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetFunctionFunctionsResponseResponseBody - OK
type GetFunctionFunctionsResponseResponseBody struct {
	// Gets a single Function.
	Data *shared.GetFunctionV1Output `json:"data,omitempty"`
}

func (o *GetFunctionFunctionsResponseResponseBody) GetData() *shared.GetFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetFunctionFunctionsResponseBody - OK
type GetFunctionFunctionsResponseBody struct {
	// Gets a single Function.
	Data *shared.GetFunctionV1Output `json:"data,omitempty"`
}

func (o *GetFunctionFunctionsResponseBody) GetData() *shared.GetFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetFunctionResponseBody - OK
type GetFunctionResponseBody struct {
	// Gets a single Function.
	Data *shared.GetFunctionV1Output `json:"data,omitempty"`
}

func (o *GetFunctionResponseBody) GetData() *shared.GetFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetFunctionResponse struct {
	// OK
	TwoHundredApplicationJSONObject *GetFunctionResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *GetFunctionFunctionsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *GetFunctionFunctionsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *GetFunctionFunctionsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetFunctionResponse) GetTwoHundredApplicationJSONObject() *GetFunctionResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetFunctionResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *GetFunctionFunctionsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *GetFunctionResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *GetFunctionFunctionsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *GetFunctionResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *GetFunctionFunctionsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *GetFunctionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFunctionResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetFunctionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFunctionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
