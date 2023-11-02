// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

// CreateLabel200ApplicationVndSegmentV1betaPlusJSON - OK
type CreateLabel200ApplicationVndSegmentV1betaPlusJSON struct {
	// Result of creating a new label in the current Workspace.
	Data *shared.CreateLabelV1Output `json:"data,omitempty"`
}

func (o *CreateLabel200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.CreateLabelV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateLabel200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateLabel200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Result of creating a new label in the current Workspace.
	Data *shared.CreateLabelV1Output `json:"data,omitempty"`
}

func (o *CreateLabel200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.CreateLabelV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateLabel200ApplicationVndSegmentV1PlusJSON - OK
type CreateLabel200ApplicationVndSegmentV1PlusJSON struct {
	// Result of creating a new label in the current Workspace.
	Data *shared.CreateLabelV1Output `json:"data,omitempty"`
}

func (o *CreateLabel200ApplicationVndSegmentV1PlusJSON) GetData() *shared.CreateLabelV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateLabel200ApplicationJSON - OK
type CreateLabel200ApplicationJSON struct {
	// Result of creating a new label in the current Workspace.
	Data *shared.CreateLabelV1Output `json:"data,omitempty"`
}

func (o *CreateLabel200ApplicationJSON) GetData() *shared.CreateLabelV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateLabelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	CreateLabel200ApplicationJSONObject *CreateLabel200ApplicationJSON
	// OK
	CreateLabel200ApplicationVndSegmentV1PlusJSONObject *CreateLabel200ApplicationVndSegmentV1PlusJSON
	// OK
	CreateLabel200ApplicationVndSegmentV1alphaPlusJSONObject *CreateLabel200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	CreateLabel200ApplicationVndSegmentV1betaPlusJSONObject *CreateLabel200ApplicationVndSegmentV1betaPlusJSON
}

func (o *CreateLabelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLabelResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateLabelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLabelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLabelResponse) GetCreateLabel200ApplicationJSONObject() *CreateLabel200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateLabel200ApplicationJSONObject
}

func (o *CreateLabelResponse) GetCreateLabel200ApplicationVndSegmentV1PlusJSONObject() *CreateLabel200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateLabel200ApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateLabelResponse) GetCreateLabel200ApplicationVndSegmentV1alphaPlusJSONObject() *CreateLabel200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateLabel200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateLabelResponse) GetCreateLabel200ApplicationVndSegmentV1betaPlusJSONObject() *CreateLabel200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateLabel200ApplicationVndSegmentV1betaPlusJSONObject
}
