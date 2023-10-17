// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type RemoveSourceConnectionFromWarehouseRequest struct {
	SourceID    string `pathParam:"style=simple,explode=false,name=sourceId"`
	WarehouseID string `pathParam:"style=simple,explode=false,name=warehouseId"`
}

// RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1betaPlusJSON - OK
type RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1betaPlusJSON struct {
	// Response indicating whether the disconnection was successful.
	Data *shared.RemoveSourceConnectionFromWarehouseV1Output `json:"data,omitempty"`
}

// RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1alphaPlusJSON - OK
type RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Response indicating whether the disconnection was successful.
	Data *shared.RemoveSourceConnectionFromWarehouseV1Output `json:"data,omitempty"`
}

// RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1PlusJSON - OK
type RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1PlusJSON struct {
	// Response indicating whether the disconnection was successful.
	Data *shared.RemoveSourceConnectionFromWarehouseV1Output `json:"data,omitempty"`
}

// RemoveSourceConnectionFromWarehouse200ApplicationJSON - OK
type RemoveSourceConnectionFromWarehouse200ApplicationJSON struct {
	// Response indicating whether the disconnection was successful.
	Data *shared.RemoveSourceConnectionFromWarehouseV1Output `json:"data,omitempty"`
}

type RemoveSourceConnectionFromWarehouseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	RemoveSourceConnectionFromWarehouse200ApplicationJSONObject *RemoveSourceConnectionFromWarehouse200ApplicationJSON
	// OK
	RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1PlusJSONObject *RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1PlusJSON
	// OK
	RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject *RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1betaPlusJSONObject *RemoveSourceConnectionFromWarehouse200ApplicationVndSegmentV1betaPlusJSON
}
