// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type CreateProfilesWarehouseRequest struct {
	CreateProfilesWarehouseAlphaInput shared.CreateProfilesWarehouseAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	SpaceID                           string                                   `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *CreateProfilesWarehouseRequest) GetCreateProfilesWarehouseAlphaInput() shared.CreateProfilesWarehouseAlphaInput {
	if o == nil {
		return shared.CreateProfilesWarehouseAlphaInput{}
	}
	return o.CreateProfilesWarehouseAlphaInput
}

func (o *CreateProfilesWarehouseRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// CreateProfilesWarehouse200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateProfilesWarehouse200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the newly created Warehouse.
	Data *shared.CreateProfilesWarehouseAlphaOutput `json:"data,omitempty"`
}

func (o *CreateProfilesWarehouse200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.CreateProfilesWarehouseAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateProfilesWarehouseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	CreateProfilesWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject *CreateProfilesWarehouse200ApplicationVndSegmentV1alphaPlusJSON
}

func (o *CreateProfilesWarehouseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateProfilesWarehouseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateProfilesWarehouseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateProfilesWarehouseResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateProfilesWarehouseResponse) GetCreateProfilesWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject() *CreateProfilesWarehouse200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateProfilesWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject
}
