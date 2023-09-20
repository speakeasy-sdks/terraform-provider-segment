// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type RemoveUsersFromUserGroupRequest struct {
	// The list of emails to remove from the user group.
	//
	// This parameter exists in v1.
	Emails      []string `queryParam:"style=deepObject,explode=true,name=emails"`
	UserGroupID string   `pathParam:"style=simple,explode=false,name=userGroupId"`
}

// RemoveUsersFromUserGroup200ApplicationVndSegmentV1betaPlusJSON - OK
type RemoveUsersFromUserGroup200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the status of the removal operation.
	Data *shared.RemoveUsersFromUserGroupV1Output `json:"data,omitempty"`
}

// RemoveUsersFromUserGroup200ApplicationVndSegmentV1alphaPlusJSON - OK
type RemoveUsersFromUserGroup200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the status of the removal operation.
	Data *shared.RemoveUsersFromUserGroupV1Output `json:"data,omitempty"`
}

// RemoveUsersFromUserGroup200ApplicationVndSegmentV1PlusJSON - OK
type RemoveUsersFromUserGroup200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the status of the removal operation.
	Data *shared.RemoveUsersFromUserGroupV1Output `json:"data,omitempty"`
}

// RemoveUsersFromUserGroup200ApplicationJSON - OK
type RemoveUsersFromUserGroup200ApplicationJSON struct {
	// Returns the status of the removal operation.
	Data *shared.RemoveUsersFromUserGroupV1Output `json:"data,omitempty"`
}

type RemoveUsersFromUserGroupResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	RemoveUsersFromUserGroup200ApplicationJSONObject *RemoveUsersFromUserGroup200ApplicationJSON
	// OK
	RemoveUsersFromUserGroup200ApplicationVndSegmentV1PlusJSONObject *RemoveUsersFromUserGroup200ApplicationVndSegmentV1PlusJSON
	// OK
	RemoveUsersFromUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject *RemoveUsersFromUserGroup200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	RemoveUsersFromUserGroup200ApplicationVndSegmentV1betaPlusJSONObject *RemoveUsersFromUserGroup200ApplicationVndSegmentV1betaPlusJSON
}