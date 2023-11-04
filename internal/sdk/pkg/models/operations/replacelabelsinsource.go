// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type ReplaceLabelsInSourceRequest struct {
	ReplaceLabelsInSourceV1Input shared.ReplaceLabelsInSourceV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	SourceID                     string                              `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *ReplaceLabelsInSourceRequest) GetReplaceLabelsInSourceV1Input() shared.ReplaceLabelsInSourceV1Input {
	if o == nil {
		return shared.ReplaceLabelsInSourceV1Input{}
	}
	return o.ReplaceLabelsInSourceV1Input
}

func (o *ReplaceLabelsInSourceRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// ReplaceLabelsInSource200ApplicationVndSegmentV1betaPlusJSON - OK
type ReplaceLabelsInSource200ApplicationVndSegmentV1betaPlusJSON struct {
	// Response from replacing a list of labels in a Source.
	Data *shared.ReplaceLabelsInSourceV1Output `json:"data,omitempty"`
}

func (o *ReplaceLabelsInSource200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.ReplaceLabelsInSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplaceLabelsInSource200ApplicationVndSegmentV1alphaPlusJSON - OK
type ReplaceLabelsInSource200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Response from replacing a list of labels in a Source.
	Data *shared.ReplaceLabelsInSourceAlphaOutput `json:"data,omitempty"`
}

func (o *ReplaceLabelsInSource200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.ReplaceLabelsInSourceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplaceLabelsInSource200ApplicationVndSegmentV1PlusJSON - OK
type ReplaceLabelsInSource200ApplicationVndSegmentV1PlusJSON struct {
	// Response from replacing a list of labels in a Source.
	Data *shared.ReplaceLabelsInSourceV1Output `json:"data,omitempty"`
}

func (o *ReplaceLabelsInSource200ApplicationVndSegmentV1PlusJSON) GetData() *shared.ReplaceLabelsInSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplaceLabelsInSource200ApplicationJSON - OK
type ReplaceLabelsInSource200ApplicationJSON struct {
	// Response from replacing a list of labels in a Source.
	Data *shared.ReplaceLabelsInSourceV1Output `json:"data,omitempty"`
}

func (o *ReplaceLabelsInSource200ApplicationJSON) GetData() *shared.ReplaceLabelsInSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ReplaceLabelsInSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ReplaceLabelsInSource200ApplicationJSONObject *ReplaceLabelsInSource200ApplicationJSON
	// OK
	ReplaceLabelsInSource200ApplicationVndSegmentV1PlusJSONObject *ReplaceLabelsInSource200ApplicationVndSegmentV1PlusJSON
	// OK
	ReplaceLabelsInSource200ApplicationVndSegmentV1alphaPlusJSONObject *ReplaceLabelsInSource200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ReplaceLabelsInSource200ApplicationVndSegmentV1betaPlusJSONObject *ReplaceLabelsInSource200ApplicationVndSegmentV1betaPlusJSON
}

func (o *ReplaceLabelsInSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReplaceLabelsInSourceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ReplaceLabelsInSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReplaceLabelsInSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReplaceLabelsInSourceResponse) GetReplaceLabelsInSource200ApplicationJSONObject() *ReplaceLabelsInSource200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ReplaceLabelsInSource200ApplicationJSONObject
}

func (o *ReplaceLabelsInSourceResponse) GetReplaceLabelsInSource200ApplicationVndSegmentV1PlusJSONObject() *ReplaceLabelsInSource200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.ReplaceLabelsInSource200ApplicationVndSegmentV1PlusJSONObject
}

func (o *ReplaceLabelsInSourceResponse) GetReplaceLabelsInSource200ApplicationVndSegmentV1alphaPlusJSONObject() *ReplaceLabelsInSource200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.ReplaceLabelsInSource200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ReplaceLabelsInSourceResponse) GetReplaceLabelsInSource200ApplicationVndSegmentV1betaPlusJSONObject() *ReplaceLabelsInSource200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.ReplaceLabelsInSource200ApplicationVndSegmentV1betaPlusJSONObject
}
