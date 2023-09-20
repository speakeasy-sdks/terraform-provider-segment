// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

// CreateUserGroup200ApplicationVndSegmentV1betaPlusJSON - OK
type CreateUserGroup200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

// CreateUserGroup200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateUserGroup200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

// CreateUserGroup200ApplicationVndSegmentV1PlusJSON - OK
type CreateUserGroup200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

// CreateUserGroup200ApplicationJSON - OK
type CreateUserGroup200ApplicationJSON struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

type CreateUserGroupResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	CreateUserGroup200ApplicationJSONObject *CreateUserGroup200ApplicationJSON
	// OK
	CreateUserGroup200ApplicationVndSegmentV1PlusJSONObject *CreateUserGroup200ApplicationVndSegmentV1PlusJSON
	// OK
	CreateUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject *CreateUserGroup200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	CreateUserGroup200ApplicationVndSegmentV1betaPlusJSONObject *CreateUserGroup200ApplicationVndSegmentV1betaPlusJSON
}
