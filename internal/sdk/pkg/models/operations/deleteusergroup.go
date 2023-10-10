// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type DeleteUserGroupRequest struct {
	UserGroupID string `pathParam:"style=simple,explode=false,name=userGroupId"`
}

// DeleteUserGroup200ApplicationVndSegmentV1betaPlusJSON - OK
type DeleteUserGroup200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the status of the completed deletion operation.
	Data *shared.DeleteUserGroupV1Output `json:"data,omitempty"`
}

// DeleteUserGroup200ApplicationVndSegmentV1alphaPlusJSON - OK
type DeleteUserGroup200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the status of the completed deletion operation.
	Data *shared.DeleteUserGroupV1Output `json:"data,omitempty"`
}

// DeleteUserGroup200ApplicationVndSegmentV1PlusJSON - OK
type DeleteUserGroup200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the status of the completed deletion operation.
	Data *shared.DeleteUserGroupV1Output `json:"data,omitempty"`
}

// DeleteUserGroup200ApplicationJSON - OK
type DeleteUserGroup200ApplicationJSON struct {
	// Returns the status of the completed deletion operation.
	Data *shared.DeleteUserGroupV1Output `json:"data,omitempty"`
}

type DeleteUserGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteUserGroup200ApplicationJSONObject *DeleteUserGroup200ApplicationJSON
	// OK
	DeleteUserGroup200ApplicationVndSegmentV1PlusJSONObject *DeleteUserGroup200ApplicationVndSegmentV1PlusJSON
	// OK
	DeleteUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject *DeleteUserGroup200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	DeleteUserGroup200ApplicationVndSegmentV1betaPlusJSONObject *DeleteUserGroup200ApplicationVndSegmentV1betaPlusJSON
}
