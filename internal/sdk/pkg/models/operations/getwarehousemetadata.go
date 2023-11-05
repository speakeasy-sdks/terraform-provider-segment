// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type GetWarehouseMetadataRequest struct {
	WarehouseMetadataID string `pathParam:"style=simple,explode=false,name=warehouseMetadataId"`
}

func (o *GetWarehouseMetadataRequest) GetWarehouseMetadataID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseMetadataID
}

// GetWarehouseMetadata200ApplicationVndSegmentV1betaPlusJSON - OK
type GetWarehouseMetadata200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a Warehouse catalog item looked up by id.
	Data *shared.GetWarehouseMetadataV1Output `json:"data,omitempty"`
}

func (o *GetWarehouseMetadata200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.GetWarehouseMetadataV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetWarehouseMetadata200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetWarehouseMetadata200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a Warehouse catalog item looked up by id.
	Data *shared.GetWarehouseMetadataV1Output `json:"data,omitempty"`
}

func (o *GetWarehouseMetadata200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.GetWarehouseMetadataV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetWarehouseMetadata200ApplicationVndSegmentV1PlusJSON - OK
type GetWarehouseMetadata200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a Warehouse catalog item looked up by id.
	Data *shared.GetWarehouseMetadataV1Output `json:"data,omitempty"`
}

func (o *GetWarehouseMetadata200ApplicationVndSegmentV1PlusJSON) GetData() *shared.GetWarehouseMetadataV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetWarehouseMetadata200ApplicationJSON - OK
type GetWarehouseMetadata200ApplicationJSON struct {
	// Returns a Warehouse catalog item looked up by id.
	Data *shared.GetWarehouseMetadataV1Output `json:"data,omitempty"`
}

func (o *GetWarehouseMetadata200ApplicationJSON) GetData() *shared.GetWarehouseMetadataV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetWarehouseMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GetWarehouseMetadata200ApplicationJSONObject *GetWarehouseMetadata200ApplicationJSON
	// OK
	GetWarehouseMetadata200ApplicationVndSegmentV1PlusJSONObject *GetWarehouseMetadata200ApplicationVndSegmentV1PlusJSON
	// OK
	GetWarehouseMetadata200ApplicationVndSegmentV1alphaPlusJSONObject *GetWarehouseMetadata200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetWarehouseMetadata200ApplicationVndSegmentV1betaPlusJSONObject *GetWarehouseMetadata200ApplicationVndSegmentV1betaPlusJSON
}

func (o *GetWarehouseMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWarehouseMetadataResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetWarehouseMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWarehouseMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWarehouseMetadataResponse) GetGetWarehouseMetadata200ApplicationJSONObject() *GetWarehouseMetadata200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetWarehouseMetadata200ApplicationJSONObject
}

func (o *GetWarehouseMetadataResponse) GetGetWarehouseMetadata200ApplicationVndSegmentV1PlusJSONObject() *GetWarehouseMetadata200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.GetWarehouseMetadata200ApplicationVndSegmentV1PlusJSONObject
}

func (o *GetWarehouseMetadataResponse) GetGetWarehouseMetadata200ApplicationVndSegmentV1alphaPlusJSONObject() *GetWarehouseMetadata200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.GetWarehouseMetadata200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *GetWarehouseMetadataResponse) GetGetWarehouseMetadata200ApplicationVndSegmentV1betaPlusJSONObject() *GetWarehouseMetadata200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.GetWarehouseMetadata200ApplicationVndSegmentV1betaPlusJSONObject
}
