// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type RemoveProfilesWarehouseFromSpaceRequest struct {
	SpaceID     string `pathParam:"style=simple,explode=false,name=spaceId"`
	WarehouseID string `pathParam:"style=simple,explode=false,name=warehouseId"`
}

func (o *RemoveProfilesWarehouseFromSpaceRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

func (o *RemoveProfilesWarehouseFromSpaceRequest) GetWarehouseID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseID
}

// RemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSON - OK
type RemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the status of a Warehouse deletion.
	Data *shared.RemoveProfilesWarehouseFromSpaceAlphaOutput `json:"data,omitempty"`
}

func (o *RemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.RemoveProfilesWarehouseFromSpaceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type RemoveProfilesWarehouseFromSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	RemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSONObject *RemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSON
}

func (o *RemoveProfilesWarehouseFromSpaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveProfilesWarehouseFromSpaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveProfilesWarehouseFromSpaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveProfilesWarehouseFromSpaceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *RemoveProfilesWarehouseFromSpaceResponse) GetRemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSONObject() *RemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.RemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSONObject
}
