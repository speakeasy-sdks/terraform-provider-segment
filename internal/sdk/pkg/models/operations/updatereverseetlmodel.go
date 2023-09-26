// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type UpdateReverseEtlModelRequest struct {
	UpdateReverseEtlModelInput shared.UpdateReverseEtlModelInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	ModelID                    string                            `pathParam:"style=simple,explode=false,name=modelId"`
}

// UpdateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON - OK
type UpdateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Defines the results of updating a Model.
	Data *shared.UpdateReverseEtlModelOutput `json:"data,omitempty"`
}

type UpdateReverseEtlModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	UpdateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSONObject *UpdateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON
}
