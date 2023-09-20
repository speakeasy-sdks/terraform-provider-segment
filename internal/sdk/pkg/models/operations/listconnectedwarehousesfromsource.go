// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type ListConnectedWarehousesFromSourceRequest struct {
	// Required pagination params for the request.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	SourceID   string                 `pathParam:"style=simple,explode=false,name=sourceId"`
}

// ListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSON - OK
type ListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a list of Warehouses connected to a Source.
	Data *shared.ListConnectedWarehousesFromSourceV1Output `json:"data,omitempty"`
}

// ListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a list of Warehouses connected to a Source.
	Data *shared.ListConnectedWarehousesFromSourceAlphaOutput `json:"data,omitempty"`
}

// ListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSON - OK
type ListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a list of Warehouses connected to a Source.
	Data *shared.ListConnectedWarehousesFromSourceV1Output `json:"data,omitempty"`
}

// ListConnectedWarehousesFromSource200ApplicationJSON - OK
type ListConnectedWarehousesFromSource200ApplicationJSON struct {
	// Returns a list of Warehouses connected to a Source.
	Data *shared.ListConnectedWarehousesFromSourceV1Output `json:"data,omitempty"`
}

type ListConnectedWarehousesFromSourceResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	ListConnectedWarehousesFromSource200ApplicationJSONObject *ListConnectedWarehousesFromSource200ApplicationJSON
	// OK
	ListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSONObject *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSON
	// OK
	ListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSONObject *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSONObject *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSON
}
