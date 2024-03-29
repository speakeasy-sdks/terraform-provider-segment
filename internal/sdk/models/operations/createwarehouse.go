// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
	"net/http"
)

// CreateWarehouseWarehousesResponse200ResponseBody - OK
type CreateWarehouseWarehousesResponse200ResponseBody struct {
	// Returns the newly created Warehouse.
	Data *shared.CreateWarehouseV1Output `json:"data,omitempty"`
}

func (o *CreateWarehouseWarehousesResponse200ResponseBody) GetData() *shared.CreateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateWarehouseWarehousesResponseResponseBody - OK
type CreateWarehouseWarehousesResponseResponseBody struct {
	// Returns the newly created Warehouse.
	Data *shared.CreateWarehouseV1Output `json:"data,omitempty"`
}

func (o *CreateWarehouseWarehousesResponseResponseBody) GetData() *shared.CreateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateWarehouseWarehousesResponseBody - OK
type CreateWarehouseWarehousesResponseBody struct {
	// Returns the newly created Warehouse.
	Data *shared.CreateWarehouseV1Output `json:"data,omitempty"`
}

func (o *CreateWarehouseWarehousesResponseBody) GetData() *shared.CreateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateWarehouseResponseBody - OK
type CreateWarehouseResponseBody struct {
	// Returns the newly created Warehouse.
	Data *shared.CreateWarehouseV1Output `json:"data,omitempty"`
}

func (o *CreateWarehouseResponseBody) GetData() *shared.CreateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateWarehouseResponse struct {
	// OK
	TwoHundredApplicationJSONObject *CreateWarehouseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *CreateWarehouseWarehousesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *CreateWarehouseWarehousesResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *CreateWarehouseWarehousesResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateWarehouseResponse) GetTwoHundredApplicationJSONObject() *CreateWarehouseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *CreateWarehouseResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *CreateWarehouseWarehousesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateWarehouseResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *CreateWarehouseWarehousesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateWarehouseResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *CreateWarehouseWarehousesResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *CreateWarehouseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateWarehouseResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateWarehouseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWarehouseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
