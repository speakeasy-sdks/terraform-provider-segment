// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListSelectiveSyncsFromWarehouseAndSpaceRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in alpha.
	Pagination  shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	SpaceID     string                 `pathParam:"style=simple,explode=false,name=spaceId"`
	WarehouseID string                 `pathParam:"style=simple,explode=false,name=warehouseId"`
}

func (o *ListSelectiveSyncsFromWarehouseAndSpaceRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

func (o *ListSelectiveSyncsFromWarehouseAndSpaceRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

func (o *ListSelectiveSyncsFromWarehouseAndSpaceRequest) GetWarehouseID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseID
}

// ListSelectiveSyncsFromWarehouseAndSpaceResponseBody - OK
type ListSelectiveSyncsFromWarehouseAndSpaceResponseBody struct {
	// Results containing the Selective Sync configuration for a Space Warehouse Connection.
	Data *shared.ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput `json:"data,omitempty"`
}

func (o *ListSelectiveSyncsFromWarehouseAndSpaceResponseBody) GetData() *shared.ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListSelectiveSyncsFromWarehouseAndSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *ListSelectiveSyncsFromWarehouseAndSpaceResponseBody
}

func (o *ListSelectiveSyncsFromWarehouseAndSpaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSelectiveSyncsFromWarehouseAndSpaceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListSelectiveSyncsFromWarehouseAndSpaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSelectiveSyncsFromWarehouseAndSpaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListSelectiveSyncsFromWarehouseAndSpaceResponse) GetObject() *ListSelectiveSyncsFromWarehouseAndSpaceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
