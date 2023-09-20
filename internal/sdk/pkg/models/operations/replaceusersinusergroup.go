// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type ReplaceUsersInUserGroupRequest struct {
	ReplaceUsersInUserGroupV1Input shared.ReplaceUsersInUserGroupV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	UserGroupID                    string                                `pathParam:"style=simple,explode=false,name=userGroupId"`
}

// ReplaceUsersInUserGroup200ApplicationVndSegmentV1betaPlusJSON - OK
type ReplaceUsersInUserGroup200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the updated user group.
	Data *shared.ReplaceUsersInUserGroupV1Output `json:"data,omitempty"`
}

// ReplaceUsersInUserGroup200ApplicationVndSegmentV1alphaPlusJSON - OK
type ReplaceUsersInUserGroup200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the updated user group.
	Data *shared.ReplaceUsersInUserGroupV1Output `json:"data,omitempty"`
}

// ReplaceUsersInUserGroup200ApplicationVndSegmentV1PlusJSON - OK
type ReplaceUsersInUserGroup200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the updated user group.
	Data *shared.ReplaceUsersInUserGroupV1Output `json:"data,omitempty"`
}

// ReplaceUsersInUserGroup200ApplicationJSON - OK
type ReplaceUsersInUserGroup200ApplicationJSON struct {
	// Returns the updated user group.
	Data *shared.ReplaceUsersInUserGroupV1Output `json:"data,omitempty"`
}

type ReplaceUsersInUserGroupResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	ReplaceUsersInUserGroup200ApplicationJSONObject *ReplaceUsersInUserGroup200ApplicationJSON
	// OK
	ReplaceUsersInUserGroup200ApplicationVndSegmentV1PlusJSONObject *ReplaceUsersInUserGroup200ApplicationVndSegmentV1PlusJSON
	// OK
	ReplaceUsersInUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject *ReplaceUsersInUserGroup200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ReplaceUsersInUserGroup200ApplicationVndSegmentV1betaPlusJSONObject *ReplaceUsersInUserGroup200ApplicationVndSegmentV1betaPlusJSON
}
