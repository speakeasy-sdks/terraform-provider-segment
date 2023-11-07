// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type RemoveAudienceFromSpaceRequest struct {
	ID      string `pathParam:"style=simple,explode=false,name=id"`
	SpaceID string `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *RemoveAudienceFromSpaceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RemoveAudienceFromSpaceRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// RemoveAudienceFromSpaceResponseBody - OK
type RemoveAudienceFromSpaceResponseBody struct {
	// Delete audience endpoint output.
	Data *shared.RemoveAudienceFromSpaceAlphaOutput `json:"data,omitempty"`
}

func (o *RemoveAudienceFromSpaceResponseBody) GetData() *shared.RemoveAudienceFromSpaceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type RemoveAudienceFromSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	Object *RemoveAudienceFromSpaceResponseBody
}

func (o *RemoveAudienceFromSpaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveAudienceFromSpaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveAudienceFromSpaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveAudienceFromSpaceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *RemoveAudienceFromSpaceResponse) GetObject() *RemoveAudienceFromSpaceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
