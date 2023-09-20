// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

// CreateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Defines the results of creating a Model.
	Data *shared.CreateReverseEtlModelOutput `json:"data,omitempty"`
}

type CreateReverseEtlModelResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	CreateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSONObject *CreateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON
}