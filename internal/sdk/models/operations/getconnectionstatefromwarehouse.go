// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
	"net/http"
)

type GetConnectionStateFromWarehouseRequest struct {
	WarehouseID string `pathParam:"style=simple,explode=false,name=warehouseId"`
}

func (o *GetConnectionStateFromWarehouseRequest) GetWarehouseID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseID
}

// GetConnectionStateFromWarehouseWarehousesResponse200ResponseBody - OK
type GetConnectionStateFromWarehouseWarehousesResponse200ResponseBody struct {
	// Returns the status of a Warehouse connection settings after an attempt to connect to it.
	Data *shared.GetConnectionStateFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *GetConnectionStateFromWarehouseWarehousesResponse200ResponseBody) GetData() *shared.GetConnectionStateFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetConnectionStateFromWarehouseWarehousesResponseResponseBody - OK
type GetConnectionStateFromWarehouseWarehousesResponseResponseBody struct {
	// Returns the status of a Warehouse connection settings after an attempt to connect to it.
	Data *shared.GetConnectionStateFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *GetConnectionStateFromWarehouseWarehousesResponseResponseBody) GetData() *shared.GetConnectionStateFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetConnectionStateFromWarehouseWarehousesResponseBody - OK
type GetConnectionStateFromWarehouseWarehousesResponseBody struct {
	// Returns the status of a Warehouse connection settings after an attempt to connect to it.
	Data *shared.GetConnectionStateFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *GetConnectionStateFromWarehouseWarehousesResponseBody) GetData() *shared.GetConnectionStateFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetConnectionStateFromWarehouseResponseBody - OK
type GetConnectionStateFromWarehouseResponseBody struct {
	// Returns the status of a Warehouse connection settings after an attempt to connect to it.
	Data *shared.GetConnectionStateFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *GetConnectionStateFromWarehouseResponseBody) GetData() *shared.GetConnectionStateFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetConnectionStateFromWarehouseResponse struct {
	// OK
	TwoHundredApplicationJSONObject *GetConnectionStateFromWarehouseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *GetConnectionStateFromWarehouseWarehousesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *GetConnectionStateFromWarehouseWarehousesResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *GetConnectionStateFromWarehouseWarehousesResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetConnectionStateFromWarehouseResponse) GetTwoHundredApplicationJSONObject() *GetConnectionStateFromWarehouseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetConnectionStateFromWarehouseResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *GetConnectionStateFromWarehouseWarehousesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *GetConnectionStateFromWarehouseResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *GetConnectionStateFromWarehouseWarehousesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *GetConnectionStateFromWarehouseResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *GetConnectionStateFromWarehouseWarehousesResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *GetConnectionStateFromWarehouseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConnectionStateFromWarehouseResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetConnectionStateFromWarehouseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectionStateFromWarehouseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
