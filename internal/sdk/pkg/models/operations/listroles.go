// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type ListRolesRequest struct {
	// Pagination for roles.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

// ListRoles200ApplicationVndSegmentV1betaPlusJSON - OK
type ListRoles200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the list of roles.
	Data *shared.ListRolesV1Output `json:"data,omitempty"`
}

// ListRoles200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListRoles200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the list of roles.
	Data *shared.ListRolesV1Output `json:"data,omitempty"`
}

// ListRoles200ApplicationVndSegmentV1PlusJSON - OK
type ListRoles200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the list of roles.
	Data *shared.ListRolesV1Output `json:"data,omitempty"`
}

// ListRoles200ApplicationJSON - OK
type ListRoles200ApplicationJSON struct {
	// Returns the list of roles.
	Data *shared.ListRolesV1Output `json:"data,omitempty"`
}

type ListRolesResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	ListRoles200ApplicationJSONObject *ListRoles200ApplicationJSON
	// OK
	ListRoles200ApplicationVndSegmentV1PlusJSONObject *ListRoles200ApplicationVndSegmentV1PlusJSON
	// OK
	ListRoles200ApplicationVndSegmentV1alphaPlusJSONObject *ListRoles200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListRoles200ApplicationVndSegmentV1betaPlusJSONObject *ListRoles200ApplicationVndSegmentV1betaPlusJSON
}
