// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type GenerateUploadURLForEdgeFunctionsRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *GenerateUploadURLForEdgeFunctionsRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// GenerateUploadURLForEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSON - OK
type GenerateUploadURLForEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Output for GenerateSignedUrl.
	Data *shared.GenerateUploadURLForEdgeFunctionsAlphaOutput `json:"data,omitempty"`
}

func (o *GenerateUploadURLForEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.GenerateUploadURLForEdgeFunctionsAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type GenerateUploadURLForEdgeFunctionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	GenerateUploadURLForEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSONObject *GenerateUploadURLForEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSON
}

func (o *GenerateUploadURLForEdgeFunctionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenerateUploadURLForEdgeFunctionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenerateUploadURLForEdgeFunctionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GenerateUploadURLForEdgeFunctionsResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GenerateUploadURLForEdgeFunctionsResponse) GetGenerateUploadURLForEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSONObject() *GenerateUploadURLForEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.GenerateUploadURLForEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSONObject
}
