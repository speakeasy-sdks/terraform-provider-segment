// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
	"net/http"
)

type ListProfilesWarehouseInSpaceRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in alpha.
	Pagination *shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	SpaceID    string                  `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *ListProfilesWarehouseInSpaceRequest) GetPagination() *shared.PaginationInput {
	if o == nil {
		return nil
	}
	return o.Pagination
}

func (o *ListProfilesWarehouseInSpaceRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// ListProfilesWarehouseInSpaceResponseBody - OK
type ListProfilesWarehouseInSpaceResponseBody struct {
	// Returns all Profiles Warehouse based on spaceID and pagination.
	Data *shared.ListProfilesWarehouseInSpaceAlphaOutput `json:"data,omitempty"`
}

func (o *ListProfilesWarehouseInSpaceResponseBody) GetData() *shared.ListProfilesWarehouseInSpaceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListProfilesWarehouseInSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *ListProfilesWarehouseInSpaceResponseBody
}

func (o *ListProfilesWarehouseInSpaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListProfilesWarehouseInSpaceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListProfilesWarehouseInSpaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListProfilesWarehouseInSpaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListProfilesWarehouseInSpaceResponse) GetObject() *ListProfilesWarehouseInSpaceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}