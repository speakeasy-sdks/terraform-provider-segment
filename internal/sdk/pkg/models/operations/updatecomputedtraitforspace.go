// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type UpdateComputedTraitForSpaceRequest struct {
	UpdateComputedTraitForSpaceAlphaInput shared.UpdateComputedTraitForSpaceAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	ID                                    string                                       `pathParam:"style=simple,explode=false,name=id"`
	SpaceID                               string                                       `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *UpdateComputedTraitForSpaceRequest) GetUpdateComputedTraitForSpaceAlphaInput() shared.UpdateComputedTraitForSpaceAlphaInput {
	if o == nil {
		return shared.UpdateComputedTraitForSpaceAlphaInput{}
	}
	return o.UpdateComputedTraitForSpaceAlphaInput
}

func (o *UpdateComputedTraitForSpaceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateComputedTraitForSpaceRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// UpdateComputedTraitForSpaceResponseBody - OK
type UpdateComputedTraitForSpaceResponseBody struct {
	// Computed Trait output for get and update.
	Data *shared.UpdateComputedTraitForSpaceAlphaOutput `json:"data,omitempty"`
}

func (o *UpdateComputedTraitForSpaceResponseBody) GetData() *shared.UpdateComputedTraitForSpaceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateComputedTraitForSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	Object *UpdateComputedTraitForSpaceResponseBody
}

func (o *UpdateComputedTraitForSpaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateComputedTraitForSpaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateComputedTraitForSpaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateComputedTraitForSpaceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *UpdateComputedTraitForSpaceResponse) GetObject() *UpdateComputedTraitForSpaceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
