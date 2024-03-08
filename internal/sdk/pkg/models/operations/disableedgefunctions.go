// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
	"net/http"
)

type DisableEdgeFunctionsRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *DisableEdgeFunctionsRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// DisableEdgeFunctionsResponseBody - OK
type DisableEdgeFunctionsResponseBody struct {
	// Output for DisableEdgeFunctions.
	Data *shared.DisableEdgeFunctionsAlphaOutput `json:"data,omitempty"`
}

func (o *DisableEdgeFunctionsResponseBody) GetData() *shared.DisableEdgeFunctionsAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type DisableEdgeFunctionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *DisableEdgeFunctionsResponseBody
}

func (o *DisableEdgeFunctionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DisableEdgeFunctionsResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *DisableEdgeFunctionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DisableEdgeFunctionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DisableEdgeFunctionsResponse) GetObject() *DisableEdgeFunctionsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
