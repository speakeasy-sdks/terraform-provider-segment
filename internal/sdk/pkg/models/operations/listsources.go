// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type ListSourcesRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

// ListSources200ApplicationVndSegmentV1betaPlusJSON - OK
type ListSources200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a list of Sources that belong to the current Workspace.
	Data *shared.ListSourcesV1Output `json:"data,omitempty"`
}

// ListSources200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListSources200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a list of Sources that belong to the current Workspace.
	Data *shared.ListSourcesAlphaOutput `json:"data,omitempty"`
}

// ListSources200ApplicationVndSegmentV1PlusJSON - OK
type ListSources200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a list of Sources that belong to the current Workspace.
	Data *shared.ListSourcesV1Output `json:"data,omitempty"`
}

// ListSources200ApplicationJSON - OK
type ListSources200ApplicationJSON struct {
	// Returns a list of Sources that belong to the current Workspace.
	Data *shared.ListSourcesV1Output `json:"data,omitempty"`
}

type ListSourcesResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	ListSources200ApplicationJSONObject *ListSources200ApplicationJSON
	// OK
	ListSources200ApplicationVndSegmentV1PlusJSONObject *ListSources200ApplicationVndSegmentV1PlusJSON
	// OK
	ListSources200ApplicationVndSegmentV1alphaPlusJSONObject *ListSources200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListSources200ApplicationVndSegmentV1betaPlusJSONObject *ListSources200ApplicationVndSegmentV1betaPlusJSON
}
