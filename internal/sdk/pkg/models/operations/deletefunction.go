// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
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

// DeleteFunction200ApplicationVndSegmentV1betaPlusJSON - OK
type DeleteFunction200ApplicationVndSegmentV1betaPlusJSON struct {
	// Delete a single Function.
	Data *shared.DeleteFunctionV1Output `json:"data,omitempty"`
}

func (o *DeleteFunction200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.DeleteFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteFunction200ApplicationVndSegmentV1alphaPlusJSON - OK
type DeleteFunction200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Delete a single Function.
	Data *shared.DeleteFunctionV1Output `json:"data,omitempty"`
}

func (o *DeleteFunction200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.DeleteFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteFunction200ApplicationVndSegmentV1PlusJSON - OK
type DeleteFunction200ApplicationVndSegmentV1PlusJSON struct {
	// Delete a single Function.
	Data *shared.DeleteFunctionV1Output `json:"data,omitempty"`
}

func (o *DeleteFunction200ApplicationVndSegmentV1PlusJSON) GetData() *shared.DeleteFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteFunction200ApplicationJSON - OK
type DeleteFunction200ApplicationJSON struct {
	// Delete a single Function.
	Data *shared.DeleteFunctionV1Output `json:"data,omitempty"`
}

func (o *DeleteFunction200ApplicationJSON) GetData() *shared.DeleteFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type DeleteFunctionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteFunction200ApplicationJSONObject *DeleteFunction200ApplicationJSON
	// OK
	DeleteFunction200ApplicationVndSegmentV1PlusJSONObject *DeleteFunction200ApplicationVndSegmentV1PlusJSON
	// OK
	DeleteFunction200ApplicationVndSegmentV1alphaPlusJSONObject *DeleteFunction200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	DeleteFunction200ApplicationVndSegmentV1betaPlusJSONObject *DeleteFunction200ApplicationVndSegmentV1betaPlusJSON
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

func (o *DeleteFunctionResponse) GetDeleteFunction200ApplicationJSONObject() *DeleteFunction200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteFunction200ApplicationJSONObject
}

func (o *DeleteFunctionResponse) GetDeleteFunction200ApplicationVndSegmentV1PlusJSONObject() *DeleteFunction200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.DeleteFunction200ApplicationVndSegmentV1PlusJSONObject
}

func (o *DeleteFunctionResponse) GetDeleteFunction200ApplicationVndSegmentV1alphaPlusJSONObject() *DeleteFunction200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.DeleteFunction200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *DeleteFunctionResponse) GetDeleteFunction200ApplicationVndSegmentV1betaPlusJSONObject() *DeleteFunction200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.DeleteFunction200ApplicationVndSegmentV1betaPlusJSONObject
}
