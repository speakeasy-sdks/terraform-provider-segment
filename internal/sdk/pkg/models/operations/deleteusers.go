// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type DeleteUsersRequest struct {
	// The ids of the users to remove.
	//
	// This parameter exists in v1.
	UserIds []string `queryParam:"style=deepObject,explode=true,name=userIds"`
}

// DeleteUsers200ApplicationVndSegmentV1betaPlusJSON - OK
type DeleteUsers200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the status of the removal operation.
	Data *shared.DeleteUsersV1Output `json:"data,omitempty"`
}

// DeleteUsers200ApplicationVndSegmentV1alphaPlusJSON - OK
type DeleteUsers200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the status of the removal operation.
	Data *shared.DeleteUsersV1Output `json:"data,omitempty"`
}

// DeleteUsers200ApplicationVndSegmentV1PlusJSON - OK
type DeleteUsers200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the status of the removal operation.
	Data *shared.DeleteUsersV1Output `json:"data,omitempty"`
}

// DeleteUsers200ApplicationJSON - OK
type DeleteUsers200ApplicationJSON struct {
	// Returns the status of the removal operation.
	Data *shared.DeleteUsersV1Output `json:"data,omitempty"`
}

type DeleteUsersResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	DeleteUsers200ApplicationJSONObject *DeleteUsers200ApplicationJSON
	// OK
	DeleteUsers200ApplicationVndSegmentV1PlusJSONObject *DeleteUsers200ApplicationVndSegmentV1PlusJSON
	// OK
	DeleteUsers200ApplicationVndSegmentV1alphaPlusJSONObject *DeleteUsers200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	DeleteUsers200ApplicationVndSegmentV1betaPlusJSONObject *DeleteUsers200ApplicationVndSegmentV1betaPlusJSON
}
