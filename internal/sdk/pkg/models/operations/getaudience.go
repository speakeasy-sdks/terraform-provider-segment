// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetAudienceRequest struct {
	ID      string `pathParam:"style=simple,explode=false,name=id"`
	SpaceID string `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *GetAudienceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetAudienceRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// GetAudienceResponseBody - OK
type GetAudienceResponseBody struct {
	// Audience output for update.
	Data *shared.GetAudienceAlphaOutput `json:"data,omitempty"`
}

func (o *GetAudienceResponseBody) GetData() *shared.GetAudienceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetAudienceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *GetAudienceResponseBody
}

func (o *GetAudienceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAudienceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetAudienceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAudienceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAudienceResponse) GetObject() *GetAudienceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
