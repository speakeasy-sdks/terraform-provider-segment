// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type ListInvitesRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

// ListInvites200ApplicationVndSegmentV1betaPlusJSON - OK
type ListInvites200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the list of invites.
	Data *shared.ListInvitesV1Output `json:"data,omitempty"`
}

// ListInvites200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListInvites200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the list of invites.
	Data *shared.ListInvitesV1Output `json:"data,omitempty"`
}

// ListInvites200ApplicationVndSegmentV1PlusJSON - OK
type ListInvites200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the list of invites.
	Data *shared.ListInvitesV1Output `json:"data,omitempty"`
}

// ListInvites200ApplicationJSON - OK
type ListInvites200ApplicationJSON struct {
	// Returns the list of invites.
	Data *shared.ListInvitesV1Output `json:"data,omitempty"`
}

type ListInvitesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ListInvites200ApplicationJSONObject *ListInvites200ApplicationJSON
	// OK
	ListInvites200ApplicationVndSegmentV1PlusJSONObject *ListInvites200ApplicationVndSegmentV1PlusJSON
	// OK
	ListInvites200ApplicationVndSegmentV1alphaPlusJSONObject *ListInvites200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListInvites200ApplicationVndSegmentV1betaPlusJSONObject *ListInvites200ApplicationVndSegmentV1betaPlusJSON
}
