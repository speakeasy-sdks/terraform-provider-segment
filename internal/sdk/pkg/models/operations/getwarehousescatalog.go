// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type GetWarehousesCatalogRequest struct {
	// Required pagination params used to filter the Warehouses catalog.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

// GetWarehousesCatalog200ApplicationVndSegmentV1betaPlusJSON - OK
type GetWarehousesCatalog200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a list of all Warehouse catalog items contained within a given page.
	Data *shared.GetWarehousesCatalogV1Output `json:"data,omitempty"`
}

// GetWarehousesCatalog200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetWarehousesCatalog200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a list of all Warehouse catalog items contained within a given page.
	Data *shared.GetWarehousesCatalogV1Output `json:"data,omitempty"`
}

// GetWarehousesCatalog200ApplicationVndSegmentV1PlusJSON - OK
type GetWarehousesCatalog200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a list of all Warehouse catalog items contained within a given page.
	Data *shared.GetWarehousesCatalogV1Output `json:"data,omitempty"`
}

// GetWarehousesCatalog200ApplicationJSON - OK
type GetWarehousesCatalog200ApplicationJSON struct {
	// Returns a list of all Warehouse catalog items contained within a given page.
	Data *shared.GetWarehousesCatalogV1Output `json:"data,omitempty"`
}

type GetWarehousesCatalogResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	GetWarehousesCatalog200ApplicationJSONObject *GetWarehousesCatalog200ApplicationJSON
	// OK
	GetWarehousesCatalog200ApplicationVndSegmentV1PlusJSONObject *GetWarehousesCatalog200ApplicationVndSegmentV1PlusJSON
	// OK
	GetWarehousesCatalog200ApplicationVndSegmentV1alphaPlusJSONObject *GetWarehousesCatalog200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetWarehousesCatalog200ApplicationVndSegmentV1betaPlusJSONObject *GetWarehousesCatalog200ApplicationVndSegmentV1betaPlusJSON
}