// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type DeleteWarehouseRequest struct {
	WarehouseID string `pathParam:"style=simple,explode=false,name=warehouseId"`
}

func (o *DeleteWarehouseRequest) GetWarehouseID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseID
}

// DeleteWarehouse200ApplicationVndSegmentV1betaPlusJSON - OK
type DeleteWarehouse200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the status of a Warehouse deletion.
	Data *shared.DeleteWarehouseV1Output `json:"data,omitempty"`
}

func (o *DeleteWarehouse200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.DeleteWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSON - OK
type DeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the status of a Warehouse deletion.
	Data *shared.DeleteWarehouseV1Output `json:"data,omitempty"`
}

func (o *DeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.DeleteWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteWarehouse200ApplicationVndSegmentV1PlusJSON - OK
type DeleteWarehouse200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the status of a Warehouse deletion.
	Data *shared.DeleteWarehouseV1Output `json:"data,omitempty"`
}

func (o *DeleteWarehouse200ApplicationVndSegmentV1PlusJSON) GetData() *shared.DeleteWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteWarehouse200ApplicationJSON - OK
type DeleteWarehouse200ApplicationJSON struct {
	// Returns the status of a Warehouse deletion.
	Data *shared.DeleteWarehouseV1Output `json:"data,omitempty"`
}

func (o *DeleteWarehouse200ApplicationJSON) GetData() *shared.DeleteWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type DeleteWarehouseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteWarehouse200ApplicationJSONObject *DeleteWarehouse200ApplicationJSON
	// OK
	DeleteWarehouse200ApplicationVndSegmentV1PlusJSONObject *DeleteWarehouse200ApplicationVndSegmentV1PlusJSON
	// OK
	DeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject *DeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	DeleteWarehouse200ApplicationVndSegmentV1betaPlusJSONObject *DeleteWarehouse200ApplicationVndSegmentV1betaPlusJSON
}

func (o *DeleteWarehouseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteWarehouseResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *DeleteWarehouseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteWarehouseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteWarehouseResponse) GetDeleteWarehouse200ApplicationJSONObject() *DeleteWarehouse200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteWarehouse200ApplicationJSONObject
}

func (o *DeleteWarehouseResponse) GetDeleteWarehouse200ApplicationVndSegmentV1PlusJSONObject() *DeleteWarehouse200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.DeleteWarehouse200ApplicationVndSegmentV1PlusJSONObject
}

func (o *DeleteWarehouseResponse) GetDeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject() *DeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.DeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *DeleteWarehouseResponse) GetDeleteWarehouse200ApplicationVndSegmentV1betaPlusJSONObject() *DeleteWarehouse200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.DeleteWarehouse200ApplicationVndSegmentV1betaPlusJSONObject
}
