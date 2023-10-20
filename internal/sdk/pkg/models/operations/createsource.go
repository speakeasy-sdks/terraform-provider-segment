// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

// CreateSource200ApplicationVndSegmentV1betaPlusJSON - OK
type CreateSource200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a newly created Source.
	Data *shared.CreateSourceV1Output `json:"data,omitempty"`
}

func (o *CreateSource200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.CreateSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateSource200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateSource200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the newly Source.
	Data *shared.CreateSourceAlphaOutput `json:"data,omitempty"`
}

func (o *CreateSource200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.CreateSourceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateSource200ApplicationVndSegmentV1PlusJSON - OK
type CreateSource200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a newly created Source.
	Data *shared.CreateSourceV1Output `json:"data,omitempty"`
}

func (o *CreateSource200ApplicationVndSegmentV1PlusJSON) GetData() *shared.CreateSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateSource200ApplicationJSON - OK
type CreateSource200ApplicationJSON struct {
	// Returns a newly created Source.
	Data *shared.CreateSourceV1Output `json:"data,omitempty"`
}

func (o *CreateSource200ApplicationJSON) GetData() *shared.CreateSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	CreateSource200ApplicationJSONObject *CreateSource200ApplicationJSON
	// OK
	CreateSource200ApplicationVndSegmentV1PlusJSONObject *CreateSource200ApplicationVndSegmentV1PlusJSON
	// OK
	CreateSource200ApplicationVndSegmentV1alphaPlusJSONObject *CreateSource200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	CreateSource200ApplicationVndSegmentV1betaPlusJSONObject *CreateSource200ApplicationVndSegmentV1betaPlusJSON
}

func (o *CreateSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSourceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSourceResponse) GetCreateSource200ApplicationJSONObject() *CreateSource200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateSource200ApplicationJSONObject
}

func (o *CreateSourceResponse) GetCreateSource200ApplicationVndSegmentV1PlusJSONObject() *CreateSource200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateSource200ApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateSourceResponse) GetCreateSource200ApplicationVndSegmentV1alphaPlusJSONObject() *CreateSource200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateSource200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateSourceResponse) GetCreateSource200ApplicationVndSegmentV1betaPlusJSONObject() *CreateSource200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateSource200ApplicationVndSegmentV1betaPlusJSONObject
}
