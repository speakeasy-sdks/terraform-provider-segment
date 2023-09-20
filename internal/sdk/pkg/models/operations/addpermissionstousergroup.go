// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type AddPermissionsToUserGroupRequest struct {
	AddPermissionsToUserGroupV1Input shared.AddPermissionsToUserGroupV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	UserGroupID                      string                                  `pathParam:"style=simple,explode=false,name=userGroupId"`
}

// AddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSON - OK
type AddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the group's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserGroupV1Output `json:"data,omitempty"`
}

// AddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSON - OK
type AddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the group's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserGroupV1Output `json:"data,omitempty"`
}

// AddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSON - OK
type AddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the group's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserGroupV1Output `json:"data,omitempty"`
}

// AddPermissionsToUserGroup200ApplicationJSON - OK
type AddPermissionsToUserGroup200ApplicationJSON struct {
	// Returns the group's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserGroupV1Output `json:"data,omitempty"`
}

type AddPermissionsToUserGroupResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	AddPermissionsToUserGroup200ApplicationJSONObject *AddPermissionsToUserGroup200ApplicationJSON
	// OK
	AddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSONObject *AddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSON
	// OK
	AddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject *AddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	AddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSONObject *AddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSON
}
