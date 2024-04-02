// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
	"net/http"
)

type UpdateWarehouseRequest struct {
	UpdateWarehouseV1Input shared.UpdateWarehouseV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	WarehouseID            string                        `pathParam:"style=simple,explode=false,name=warehouseId"`
}

func (o *UpdateWarehouseRequest) GetUpdateWarehouseV1Input() shared.UpdateWarehouseV1Input {
	if o == nil {
		return shared.UpdateWarehouseV1Input{}
	}
	return o.UpdateWarehouseV1Input
}

func (o *UpdateWarehouseRequest) GetWarehouseID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseID
}

// UpdateWarehouseWarehousesResponse200ResponseBody - OK
type UpdateWarehouseWarehousesResponse200ResponseBody struct {
	// Returns the updated Warehouse.
	Data *shared.UpdateWarehouseV1Output `json:"data,omitempty"`
}

func (o *UpdateWarehouseWarehousesResponse200ResponseBody) GetData() *shared.UpdateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateWarehouseWarehousesResponseResponseBody - OK
type UpdateWarehouseWarehousesResponseResponseBody struct {
	// Returns the updated Warehouse.
	Data *shared.UpdateWarehouseV1Output `json:"data,omitempty"`
}

func (o *UpdateWarehouseWarehousesResponseResponseBody) GetData() *shared.UpdateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateWarehouseWarehousesResponseBody - OK
type UpdateWarehouseWarehousesResponseBody struct {
	// Returns the updated Warehouse.
	Data *shared.UpdateWarehouseV1Output `json:"data,omitempty"`
}

func (o *UpdateWarehouseWarehousesResponseBody) GetData() *shared.UpdateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateWarehouseResponseBody - OK
type UpdateWarehouseResponseBody struct {
	// Returns the updated Warehouse.
	Data *shared.UpdateWarehouseV1Output `json:"data,omitempty"`
}

func (o *UpdateWarehouseResponseBody) GetData() *shared.UpdateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateWarehouseResponse struct {
	// OK
	TwoHundredApplicationJSONObject *UpdateWarehouseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *UpdateWarehouseWarehousesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *UpdateWarehouseWarehousesResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *UpdateWarehouseWarehousesResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateWarehouseResponse) GetTwoHundredApplicationJSONObject() *UpdateWarehouseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *UpdateWarehouseResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *UpdateWarehouseWarehousesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *UpdateWarehouseResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *UpdateWarehouseWarehousesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *UpdateWarehouseResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *UpdateWarehouseWarehousesResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *UpdateWarehouseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateWarehouseResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *UpdateWarehouseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateWarehouseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}