// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type GetReverseEtlModelRequest struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

// GetReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Defines the result of getting a Model.
	Data *shared.GetReverseEtlModelOutput `json:"data,omitempty"`
}

type GetReverseEtlModelResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	GetReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSONObject *GetReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON
}