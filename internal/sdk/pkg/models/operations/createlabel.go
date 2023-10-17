// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

// CreateLabel200ApplicationVndSegmentV1betaPlusJSON - OK
type CreateLabel200ApplicationVndSegmentV1betaPlusJSON struct {
	// Result of creating a new label in the current Workspace.
	Data *shared.CreateLabelV1Output `json:"data,omitempty"`
}

// CreateLabel200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateLabel200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Result of creating a new label in the current Workspace.
	Data *shared.CreateLabelV1Output `json:"data,omitempty"`
}

// CreateLabel200ApplicationVndSegmentV1PlusJSON - OK
type CreateLabel200ApplicationVndSegmentV1PlusJSON struct {
	// Result of creating a new label in the current Workspace.
	Data *shared.CreateLabelV1Output `json:"data,omitempty"`
}

// CreateLabel200ApplicationJSON - OK
type CreateLabel200ApplicationJSON struct {
	// Result of creating a new label in the current Workspace.
	Data *shared.CreateLabelV1Output `json:"data,omitempty"`
}

type CreateLabelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	CreateLabel200ApplicationJSONObject *CreateLabel200ApplicationJSON
	// OK
	CreateLabel200ApplicationVndSegmentV1PlusJSONObject *CreateLabel200ApplicationVndSegmentV1PlusJSON
	// OK
	CreateLabel200ApplicationVndSegmentV1alphaPlusJSONObject *CreateLabel200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	CreateLabel200ApplicationVndSegmentV1betaPlusJSONObject *CreateLabel200ApplicationVndSegmentV1betaPlusJSON
}
