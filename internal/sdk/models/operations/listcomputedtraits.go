// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
	"net/http"
)

type ListComputedTraitsRequest struct {
	// Information about the pagination of this response.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	SpaceID    string                 `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *ListComputedTraitsRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

func (o *ListComputedTraitsRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// ListComputedTraitsResponseBody - OK
type ListComputedTraitsResponseBody struct {
	// List computed traits endpoint output.
	Data *shared.ListComputedTraitsAlphaOutput `json:"data,omitempty"`
}

func (o *ListComputedTraitsResponseBody) GetData() *shared.ListComputedTraitsAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListComputedTraitsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *ListComputedTraitsResponseBody
}

func (o *ListComputedTraitsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListComputedTraitsResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListComputedTraitsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListComputedTraitsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListComputedTraitsResponse) GetObject() *ListComputedTraitsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}