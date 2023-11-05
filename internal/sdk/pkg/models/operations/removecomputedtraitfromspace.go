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

func (o *RemoveComputedTraitFromSpaceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RemoveComputedTraitFromSpaceRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// RemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSON - OK
type RemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Delete computed trait endpoint output.
	Data *shared.RemoveComputedTraitFromSpaceAlphaOutput `json:"data,omitempty"`
}

func (o *RemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.RemoveComputedTraitFromSpaceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type RemoveComputedTraitFromSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	RemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSONObject *RemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSON
}

func (o *RemoveComputedTraitFromSpaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveComputedTraitFromSpaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveComputedTraitFromSpaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveComputedTraitFromSpaceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *RemoveComputedTraitFromSpaceResponse) GetRemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSONObject() *RemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.RemoveComputedTraitFromSpace200ApplicationVndSegmentV1alphaPlusJSONObject
}
