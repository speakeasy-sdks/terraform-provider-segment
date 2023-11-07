// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type ListConnectedSourcesFromWarehouseRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in v1.
	Pagination  shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	WarehouseID string                 `pathParam:"style=simple,explode=false,name=warehouseId"`
}

func (o *ListConnectedSourcesFromWarehouseRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

func (o *ListConnectedSourcesFromWarehouseRequest) GetWarehouseID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseID
}

// ListConnectedSourcesFromWarehouseWarehousesResponse200ResponseBody - OK
type ListConnectedSourcesFromWarehouseWarehousesResponse200ResponseBody struct {
	// Returns a list of Sources connected to a Warehouse.
	Data *shared.ListConnectedSourcesFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *ListConnectedSourcesFromWarehouseWarehousesResponse200ResponseBody) GetData() *shared.ListConnectedSourcesFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListConnectedSourcesFromWarehouseWarehousesResponseResponseBody - OK
type ListConnectedSourcesFromWarehouseWarehousesResponseResponseBody struct {
	// Returns a list of Sources connected to a Warehouse.
	Data *shared.ListConnectedSourcesFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *ListConnectedSourcesFromWarehouseWarehousesResponseResponseBody) GetData() *shared.ListConnectedSourcesFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListConnectedSourcesFromWarehouseWarehousesResponseBody - OK
type ListConnectedSourcesFromWarehouseWarehousesResponseBody struct {
	// Returns a list of Sources connected to a Warehouse.
	Data *shared.ListConnectedSourcesFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *ListConnectedSourcesFromWarehouseWarehousesResponseBody) GetData() *shared.ListConnectedSourcesFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListConnectedSourcesFromWarehouseResponseBody - OK
type ListConnectedSourcesFromWarehouseResponseBody struct {
	// Returns a list of Sources connected to a Warehouse.
	Data *shared.ListConnectedSourcesFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *ListConnectedSourcesFromWarehouseResponseBody) GetData() *shared.ListConnectedSourcesFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListConnectedSourcesFromWarehouseResponse struct {
	// OK
	TwoHundredApplicationJSONObject *ListConnectedSourcesFromWarehouseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *ListConnectedSourcesFromWarehouseWarehousesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *ListConnectedSourcesFromWarehouseWarehousesResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *ListConnectedSourcesFromWarehouseWarehousesResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListConnectedSourcesFromWarehouseResponse) GetTwoHundredApplicationJSONObject() *ListConnectedSourcesFromWarehouseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ListConnectedSourcesFromWarehouseResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *ListConnectedSourcesFromWarehouseWarehousesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *ListConnectedSourcesFromWarehouseResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *ListConnectedSourcesFromWarehouseWarehousesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ListConnectedSourcesFromWarehouseResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *ListConnectedSourcesFromWarehouseWarehousesResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *ListConnectedSourcesFromWarehouseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListConnectedSourcesFromWarehouseResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListConnectedSourcesFromWarehouseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListConnectedSourcesFromWarehouseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
