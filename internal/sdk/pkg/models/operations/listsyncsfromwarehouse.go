// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type ListSyncsFromWarehouseRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in v1.
	Pagination  shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	WarehouseID string                 `pathParam:"style=simple,explode=false,name=warehouseId"`
}

// ListSyncsFromWarehouse200ApplicationVndSegmentV1betaPlusJSON - OK
type ListSyncsFromWarehouse200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns an overview page that contains the last syncs for a Warehouse.
	Data *shared.ListSyncsFromWarehouseV1Output `json:"data,omitempty"`
}

// ListSyncsFromWarehouse200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListSyncsFromWarehouse200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns an overview page that contains the last syncs for a Warehouse.
	Data *shared.ListSyncsFromWarehouseV1Output `json:"data,omitempty"`
}

// ListSyncsFromWarehouse200ApplicationVndSegmentV1PlusJSON - OK
type ListSyncsFromWarehouse200ApplicationVndSegmentV1PlusJSON struct {
	// Returns an overview page that contains the last syncs for a Warehouse.
	Data *shared.ListSyncsFromWarehouseV1Output `json:"data,omitempty"`
}

// ListSyncsFromWarehouse200ApplicationJSON - OK
type ListSyncsFromWarehouse200ApplicationJSON struct {
	// Returns an overview page that contains the last syncs for a Warehouse.
	Data *shared.ListSyncsFromWarehouseV1Output `json:"data,omitempty"`
}

type ListSyncsFromWarehouseResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	ListSyncsFromWarehouse200ApplicationJSONObject *ListSyncsFromWarehouse200ApplicationJSON
	// OK
	ListSyncsFromWarehouse200ApplicationVndSegmentV1PlusJSONObject *ListSyncsFromWarehouse200ApplicationVndSegmentV1PlusJSON
	// OK
	ListSyncsFromWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject *ListSyncsFromWarehouse200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListSyncsFromWarehouse200ApplicationVndSegmentV1betaPlusJSONObject *ListSyncsFromWarehouse200ApplicationVndSegmentV1betaPlusJSON
}