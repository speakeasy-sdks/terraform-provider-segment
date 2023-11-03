// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type CreateFunctionDeploymentRequest struct {
	FunctionID string `pathParam:"style=simple,explode=false,name=functionId"`
}

func (o *CreateFunctionDeploymentRequest) GetFunctionID() string {
	if o == nil {
		return ""
	}
	return o.FunctionID
}

// CreateFunctionDeployment200ApplicationVndSegmentV1betaPlusJSON - OK
type CreateFunctionDeployment200ApplicationVndSegmentV1betaPlusJSON struct {
	// Updates the deployment for a Source Function instance.
	Data *shared.CreateFunctionDeploymentV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionDeployment200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.CreateFunctionDeploymentV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateFunctionDeployment200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateFunctionDeployment200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Updates the deployment for a Source Function instance.
	Data *shared.CreateFunctionDeploymentV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionDeployment200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.CreateFunctionDeploymentV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateFunctionDeployment200ApplicationVndSegmentV1PlusJSON - OK
type CreateFunctionDeployment200ApplicationVndSegmentV1PlusJSON struct {
	// Updates the deployment for a Source Function instance.
	Data *shared.CreateFunctionDeploymentV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionDeployment200ApplicationVndSegmentV1PlusJSON) GetData() *shared.CreateFunctionDeploymentV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateFunctionDeployment200ApplicationJSON - OK
type CreateFunctionDeployment200ApplicationJSON struct {
	// Updates the deployment for a Source Function instance.
	Data *shared.CreateFunctionDeploymentV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionDeployment200ApplicationJSON) GetData() *shared.CreateFunctionDeploymentV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateFunctionDeploymentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	CreateFunctionDeployment200ApplicationJSONObject *CreateFunctionDeployment200ApplicationJSON
	// OK
	CreateFunctionDeployment200ApplicationVndSegmentV1PlusJSONObject *CreateFunctionDeployment200ApplicationVndSegmentV1PlusJSON
	// OK
	CreateFunctionDeployment200ApplicationVndSegmentV1alphaPlusJSONObject *CreateFunctionDeployment200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	CreateFunctionDeployment200ApplicationVndSegmentV1betaPlusJSONObject *CreateFunctionDeployment200ApplicationVndSegmentV1betaPlusJSON
}

func (o *CreateFunctionDeploymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateFunctionDeploymentResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateFunctionDeploymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateFunctionDeploymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateFunctionDeploymentResponse) GetCreateFunctionDeployment200ApplicationJSONObject() *CreateFunctionDeployment200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateFunctionDeployment200ApplicationJSONObject
}

func (o *CreateFunctionDeploymentResponse) GetCreateFunctionDeployment200ApplicationVndSegmentV1PlusJSONObject() *CreateFunctionDeployment200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateFunctionDeployment200ApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateFunctionDeploymentResponse) GetCreateFunctionDeployment200ApplicationVndSegmentV1alphaPlusJSONObject() *CreateFunctionDeployment200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateFunctionDeployment200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateFunctionDeploymentResponse) GetCreateFunctionDeployment200ApplicationVndSegmentV1betaPlusJSONObject() *CreateFunctionDeployment200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateFunctionDeployment200ApplicationVndSegmentV1betaPlusJSONObject
}
