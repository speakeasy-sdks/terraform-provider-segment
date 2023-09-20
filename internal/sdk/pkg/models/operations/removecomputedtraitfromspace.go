// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type RemoveComputedTraitFromSpaceRequest struct {
	ID      string `pathParam:"style=simple,explode=false,name=id"`
	SpaceID string `pathParam:"style=simple,explode=false,name=spaceId"`
}

// RemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSON - OK
type RemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Delete computed trait endpoint output.
	Data *shared.RemoveComputedTraitFromSpaceAlphaOutput `json:"data,omitempty"`
}

type RemoveComputedTraitFromSpaceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	RemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSONObject *RemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSON
}
